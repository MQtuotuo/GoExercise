package main

import (
	"fmt"
)

type XmlParser struct {
	Content   string
	CreatedAt string
}

func (x XmlParser) Parse() string {
	return fmt.Sprintf("Parsing the content in XML: %s", x.Content)
}

func (x XmlParser) FormatDate() string {
	return fmt.Sprintf("Formatting date for XML: %s", x.CreatedAt)
}

type JsonParser struct {
	Content   string
	CreatedAt string
}

func (j JsonParser) Parse() string {
	return fmt.Sprintf("Parsing the content in JSON: %s", j.Content)
}

func (j JsonParser) FormatDate() string {
	return fmt.Sprintf("Formatting date for JSON: %s", j.CreatedAt)
}

type CsvParser struct {
	Content   string
	CreatedAt string
}

func (c CsvParser) Parse() string {
	return fmt.Sprintf("Parsing the content in CSV: %s", c.Content)
}

func (c CsvParser) FormatDate() string {
	return fmt.Sprintf("Formatting date for CSV: %s", c.CreatedAt)
}

func XmlPrinter(xml XmlParser) string {
	return xml.Parse() + " - Date - " + xml.FormatDate()
}

func JsonPrinter(json JsonParser) string {
	return json.Parse() + " - Date - " + json.FormatDate()
}

func CsvPrinter(csv CsvParser) string {
	return csv.Parse() + " - Date - " + csv.FormatDate()
}

type AvroParser struct {
	Content   string
	CreatedAt string
}

func (a AvroParser) Parse() string {
	return fmt.Sprintf("Parsing the content in Avro: %s", a.Content)
}

func (a AvroParser) FormatDate() string {
	return fmt.Sprintf("Formatting date for Avro: %s", a.CreatedAt)
}

type Parser interface {
	Parse() string
	FormatDate() string
}

func Printer(format Parser) string {
	return format.Parse() + " - Date - " + format.FormatDate()
}

func main() {
	xmlParser := XmlParser{Content: "<title>My Title<title>", CreatedAt: "2018-01-20"}
	xmlPrinter := XmlPrinter(xmlParser)
	fmt.Println(xmlPrinter)
	fmt.Println(Printer(xmlParser))

	jsonParser := JsonParser{Content: "{'title': 'My Title'}", CreatedAt: "2018-01-25"}
	jsonPrinter := JsonPrinter(jsonParser)
	fmt.Println(jsonPrinter)
	fmt.Println(Printer(jsonParser))

	csvParser := CsvParser{Content: "My Content; Another Content", CreatedAt: "2018-01-25"}
	csvPrinter := CsvPrinter(csvParser)
	fmt.Println(csvPrinter)
	fmt.Println(Printer(csvParser))

	avroParser := AvroParser{Content: "My avro content", CreatedAt: "2018-01-25"}
	avroPrinter := Printer(avroParser)
	fmt.Println(avroPrinter)

}
