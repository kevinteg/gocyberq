package gocyberq

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
  // "fmt"
)

var _ = Describe("Cyberq", func() {
  It("should have a test", func() {
    expectedValue := Nutcstatus{
      OutputPercent: "100",
      TimerCurrent: "00:00:00",
      CookTemp: "3343",
      Food1Temp: "823",
      Food2Temp: "OPEN",
      Food3Temp: "OPEN",
      CookStatus: "0",
      Food1Status: "0",
      Food2Status: "4",
      Food3Status: "4",
      TimerStatus: "0",
      DegUnits: "1",
      CookCycleTime: "6",
      CookPropBand: "500",
      CookRamp: "0",
    }
    Expect(Status()).To(Equal(expectedValue))
  })

  It("should return a all object", func() {
    // status := All()
    // fmt.Println(status)
    expectedValue := Nutcallstatus{
      Cook: Cook{
        Name: "Big Green Egg",
        Temp: "3216",
        Set: "4000",
        Status: "0",
      },
      Food1: Food1{
        Name: "Chicken Quarters",
        Temp: "1482",
        Set: "1750",
        Status: "0",
      },
      Food2: Food2{
        Name: "Food2",
        Temp: "OPEN",
        Set: "1000",
        Status: "4",
      },
      Food3: Food3{
        Name: "Food3",
        Temp: "OPEN",
        Set: "1000",
        Status: "4",
      },
      OutputPercent: "100",
      TimerCurrent: "00:00:00",
      TimerStatus: "0",
      DegUnits: "1",
      CookCycleTime: "6",
      CookPropBand: "500",
      CookRamp: "0",
    }
    Expect(All()).To(Equal(expectedValue))
  })

  It("should return a config object", func() {
    // status := All()
    // fmt.Println(status)
    expectedValue := Nutcallstatus{
      Cook: Cook{
        Name: "Big Green Egg",
        Temp: "3220",
        Set: "4000",
        Status: "0",
      },
      Food1: Food1{
        Name: "Chicken Quarters",
        Temp: "1493",
        Set: "1750",
        Status: "0",
      },
      Food2: Food2{
        Name: "Food2",
        Temp: "OPEN",
        Set: "1000",
        Status: "4",
      },
      Food3: Food3{
        Name: "Food3",
        Temp: "OPEN",
        Set: "1000",
        Status: "4",
      },
      OutputPercent: "100",
      TimerCurrent: "00:00:00",
      TimerStatus: "0",
      DegUnits: "",
      CookCycleTime: "",
      CookPropBand: "",
      CookRamp: "",
      System: System{
            MenuScrolling: "1",
            LcdBacklight: "47",
            LcdContrast: "10",
            DegUnits: "1",
            AlarmBeeps: "0",
            KeyBeeps: "0",
      },
      Control: Control{
            TimeoutAction: "0",
            CookHold: "2000",
            AlarmDev: "500",
            CookRamp: "0",
            OpenDetect: "1",
            CycleTime: "6",
            PropBand: "500",
      },
    }
    actualValue := Config()
    Expect(actualValue.Cook).To(Equal(expectedValue.Cook))
    Expect(actualValue.Food1).To(Equal(expectedValue.Food1))
    Expect(actualValue.Food2).To(Equal(expectedValue.Food2))
    Expect(actualValue.Food3).To(Equal(expectedValue.Food3))
    Expect(actualValue.System).To(Equal(expectedValue.System))
    Expect(actualValue.Control).To(Equal(expectedValue.Control))
    Expect(actualValue).To(Equal(expectedValue))
  })
})
