package decisions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func AppendLocalDecision(decision Decision) {
	jsonFile, err := os.OpenFile("declog/declog.json", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var decisions Decisions

	err = json.Unmarshal(byteValue, &decisions)
	if err != nil {
		log.Fatal(err)
	}

	decisions.Decisions = append(decisions.Decisions, decision)

	jsonFile.Truncate(0)
	jsonFile.Seek(0, 0)

	newDecisionsJson, err := json.MarshalIndent(decisions, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	jsonFile.Write(newDecisionsJson)
}
