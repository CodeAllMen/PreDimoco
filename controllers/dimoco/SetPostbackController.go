package dimoco

import (
	"encoding/json"
	"fmt"
	"github.com/MobileCPX/PreBaseLib/common"
	"github.com/MobileCPX/PreDimoco/models/dimoco"
	"io/ioutil"
)

type SetPostbackController struct {
	common.BaseController
}

func (c *SetPostbackController) Get() {
	affName := c.GetString("aff_name")
	promoter := c.GetString("promoter")
	postbackURL := c.GetString("postback_url")
	payout, _ := c.GetFloat("payout")
	postbackRate, _ := c.GetInt("rate")
	offerID, _ := c.GetInt("offer_id")
	campID, _ := c.GetInt("camp_id")
	fmt.Println(affName, promoter, postbackURL, offerID, campID)
	postback := new(dimoco.Postback)
	if offerID != 0 && campID != 0 {
		err := postback.CheckOfferID(int64(offerID))
		if err == nil && postback.ID != 0 {
			fmt.Println("ERROR,OfferID已经存在")
			c.Ctx.WriteString("ERROR,OfferID已经存在")
			c.StopRun()
		} else {
			postback.AffName = affName
			postback.PromoterName = promoter
			postback.OfferID = offerID
			postback.CampID = campID
			postback.PostbackRate = postbackRate
			postback.PostbackURL = postbackURL
			postback.Payout = float32(payout)
			if postbackRate == 0 {
				postbackRate = 50
			}
			err = postback.InsertPostback()
			if err != nil {
				fmt.Println("ERROR,插入postbak失败")
				c.Ctx.WriteString("ERROR,插入postbak失败")
				c.StopRun()
			}
		}
	} else {
		fmt.Println("ERROR,offerID 是空")
		c.Ctx.WriteString("ERROR,offerID 是空")
		c.StopRun()
	}
	c.Ctx.WriteString("SUCCESS")
}

func (c *SetPostbackController) Post() {
	postback := new(dimoco.Postback)

	reqBody := c.Ctx.Request.Body
	reqByte, err := ioutil.ReadAll(reqBody)
	if err == nil {
		_ = json.Unmarshal(reqByte, postback)
		fmt.Println(postback)
	} else {
		c.StringResult("ERROR,json解析失败： " + err.Error())
	}

	if postback.OfferID != 0 && postback.CampID != 0 {
		err := postback.CheckOfferID(int64(postback.OfferID))
		if err == nil && postback.ID != 0 {
			fmt.Println("ERROR,OfferID已经存在")
			c.Ctx.WriteString("ERROR,OfferID已经存在")
			c.StopRun()
		} else {
			err = postback.InsertPostback()
			if err != nil {
				fmt.Println("ERROR,插入postbak失败")
				c.Ctx.WriteString("ERROR,插入postbak失败")
				c.StopRun()
			}
		}
	} else {
		fmt.Println("ERROR,offerID 是空")
		c.Ctx.WriteString("ERROR,offerID 是空")
		c.StopRun()
	}
	c.Ctx.WriteString("SUCCESS")
}
