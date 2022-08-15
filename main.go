package main

import (
	"golang.org/x/sys/windows/registry"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	for {
		_, alreadyExisted, err := registry.CreateKey(registry.LOCAL_MACHINE, `SYSTEM\Setup\LabConfig`, registry.ALL_ACCESS)
		if err != nil {
			log.Fatal(err)
		}
		if alreadyExisted {
			err := registry.DeleteKey(registry.LOCAL_MACHINE, `SYSTEM\Setup\LabConfig`)
			if err != nil {
				log.Fatal(err)
			}
			continue
		} else {
			break
		}
	}
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\Setup\LabConfig`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer func(k registry.Key) {
		err := k.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(k)

	keys := []string{"BypassCPUCheck", "BypassStorageCheck", "BypassRAMCheck", "BypassTPMCheck", "BypassSecureBootCheck"}
	for _, key := range keys {
		err = k.SetDWordValue(key, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
