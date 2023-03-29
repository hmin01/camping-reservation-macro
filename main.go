package main

import (
	"log"
	"strconv"
	"time"

	// Web driver
	"github.com/fedesog/webdriver"
	// Selenium
	"github.com/tebeka/selenium"
)

func main() {
	// 드라이버 생성
	chromDriver := webdriver.NewChromeDriver("./chromedriver.exe")
	// 드라이버 시작
	err := chromDriver.Start()
	// 예외 처리
	if err != nil {
		log.Fatalln(err)
	}

	desired := webdriver.Capabilities{"Platform": "Windows"}
	required := webdriver.Capabilities{}
	// 세션 생성
	session, err := chromDriver.NewSession(desired, required)
	// 예외 처리
	if err != nil {
		log.Fatalln(err)
	}

	// URL 접근
	err = session.Url("https://camping.gtdc.or.kr/DZ_reservation/reserCamping_v3.php?xch=reservation&xid=camping_reservation&sdate=202304")
	// 예외 처리
	if err != nil {
		log.Fatalln(err)
	}

	// 동의 모달 버튼
	agreeBtn, err := session.FindElement(selenium.ByCSSSelector, ".denoPopupBox > .contentArea button")
	// 예외 처리
	if err != nil {
		log.Fatalln(err)
	}
	agreeBtn.Click()
	// 대기
	time.Sleep(10)

	// 예약 버튼
	reservationBtn, err := session.FindElement(selenium.ByCSSSelector, "div.reservationZone table.dztbl > tbody > tr:nth-child(6) > td:first-child > ul > li:last-child > button")
	// 예외 처리
	if err != nil {
		log.Fatalln(err)
	}
	reservationBtn.Click()
	// 대기
	time.Sleep(200 * time.Millisecond)

	// 캠핑 위치 div
	location, err := session.FindElement(selenium.ByCSSSelector, "#camping_zone")
	// 캠핑 위치 버튼
	var locationBtn webdriver.WebElement
	for i := 3; i <= 10; i++ {
		// 캠핑 위치 버튼
		locationBtn, err = location.FindElement(selenium.ByCSSSelector, "#camping_zone > button:nth-child("+strconv.Itoa(i)+").on")
		// 예외 처리
		if err != nil {
			if i == 10 {
				log.Fatalln("[ERROR] Not found an element")
				return
			}
		} else {
			break
		}
	}
	locationBtn.Click()
	// 대기
	time.Sleep(10)

	// 인원 선택 셀럭트
	countSelect, err := session.FindElement(selenium.ByCSSSelector, "div.reservationZone table.setPersion_tbl > tbody > tr > td:nth-child(4) > select")
	// 예외 처리
	if err != nil {
		log.Fatalln(err)
	}
	countSelect.SendKeys("2")

	// 기간 선택 셀럭트
	periodSelect, err := session.FindElement(selenium.ByCSSSelector, "select#reservation_period")
	// 예외 처리
	if err != nil {
		log.Fatalln(err)
	}
	periodSelect.SendKeys("2")

	// 매크로 방지 입력 창
	mnInput, err := session.FindElement(selenium.ByCSSSelector, "input#CAPTCHA_TEXT")
	// 예외 처리
	if err != nil {
		log.Fatalln(err)
	}
	mnInput.Click()

	// // 세션 종료
	// session.Delete()
	// // 웹 드라이버 종료
	// chromDriver.Stop()
}
