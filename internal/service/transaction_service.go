package service

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Djuanzz/cashlens-backend/internal/model"
	"github.com/ledongthuc/pdf"
)

type TransactionService struct {
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (s *TransactionService) ProcessTransactionFile(filePath string) error {

	text, err := extractTextFromPDF(filePath)
	if err != nil {
		return err
	}

	cleaned := cleanExtractedText(text)
	lines := strings.Split(cleaned, "\n")
	var results []model.Transaction

	for _, line := range lines {
		txn := parseTransaction(line)
		if txn != nil {
			results = append(results, *txn)
		}
	}

	saveJSON("transactions.json", results)
	return nil
}

func (s *TransactionService) BlueTakeJson() ([]model.Transaction, error) {
	file, err := os.ReadFile("files/transactions.json")

	if err != nil {
		return nil, err
	}

	var transactions []model.Transaction
	err = json.Unmarshal(file, &transactions)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func extractTextFromPDF(filePath string) (string, error) {
	f, r, err := pdf.Open(filePath)
	if err != nil {
		return "", err
	}

	defer f.Close()

	var text string

	for i := 1; i <= r.NumPage(); i++ {
		page := r.Page(i)
		if page.V.IsNull() {
			continue
		}
		content, _ := page.GetPlainText(nil)
		text += content
	}

	return text, nil
}

func cleanExtractedText(text string) string {
	text = strings.ReplaceAll(text, "\r", "")
	lines := strings.Split(text, "\n")
	var result []string
	var current string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if isNoise(line) {
			continue
		}
		if isDateLine(line) {
			if current != "" {
				result = append(result, current)
			}
			current = line
			continue
		}
		if current != "" {
			current += " " + line
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return strings.Join(result, "\n")
}

func isNoise(line string) bool {
	noisePatterns := []string{
		"SALDO AWAL",
		"SALDO AKHIR",
		"MUTASI CR",
		"MUTASI DB",
		"Bersambung",
		"REKENING TAHAPAN",
		"KCU",
		"NO. REKENING",
		"HALAMAN",
		"PERIODE",
		"MATA UANG",
		"CATATAN",
		"KETERANGAN",
		"CBG",
		"MUTASI SALDO",
		"INDONESIA",
	}

	for _, pattern := range noisePatterns {
		if strings.Contains(line, pattern) {
			return true
		}
	}
	return false
}

func isDateLine(line string) bool {
	return regexp.MustCompile(`^\d{2}/\d{2}`).MatchString(line)
}

var amountRegex = regexp.MustCompile(`\d{1,3}(?:,\d{3})+\.\d{2}`)

func parseAmount(raw string) float64 {
	cleaned := strings.ReplaceAll(raw, ",", "")
	val, _ := strconv.ParseFloat(cleaned, 64)
	return val
}

func extractAmount(line string) (float64, string, bool) {
	reDB := regexp.MustCompile(`(\d{1,3}(?:,\d{3})+\.\d{2})\s+DB`)
	match := reDB.FindStringSubmatch(line)
	if len(match) > 1 {
		return parseAmount(match[1]), "DB", true
	}
	reAmount := regexp.MustCompile(`\d{1,3}(?:,\d{3})+\.\d{2}`)
	matches := reAmount.FindAllString(line, -1)
	if len(matches) == 0 {
		return 0, "", false
	}

	return parseAmount(matches[0]), "CR", true

}

func parseTransaction(line string) *model.Transaction {
	line = strings.TrimSpace(line)
	if len(line) < 10 {
		return nil
	}
	matches := amountRegex.FindAllString(line, -1)
	if len(matches) == 0 {
		return nil
	}
	rawAmount := matches[0]
	amount, typ, ok := extractAmount(line)
	if !ok {
		return nil
	}
	if strings.Contains(line, "DB") {
		typ = "DB"
	}
	date := line[:5]
	idx := strings.Index(line, rawAmount)
	if idx != -1 {
		line = line[:idx+len(rawAmount)]
	}
	desc := line
	desc = strings.Replace(desc, rawAmount, "", 1)
	desc = strings.Replace(desc, date, "", 1)
	noise := []string{
		"DB", "CR",
		"TGL:", "TANGGAL",
		"WSID", "FTSCY",
		"BIF", "TRANSFER",
		"DR", "TRSF",
	}

	for _, n := range noise {
		desc = strings.ReplaceAll(desc, n, "")
	}
	desc = strings.Join(strings.Fields(desc), " ")

	return &model.Transaction{
		Date:   date,
		Amount: amount,
		Type:   typ,
		Desc:   desc,
	}
}

func saveJSON(path string, data []model.Transaction) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
