package dimoco

import (
	"fmt"
	"github.com/MobileCPX/PreDimoco/models/dimoco"
	"github.com/astaxie/beego"
)

type SetPostbackController struct {
	beego.Controller
}

func (c *SetPostbackController) Get() {
	affName := c.GetString("aff_name")
	promoter := c.GetString("promoter")
	postbackURL := c.GetString("postback_url")
	payout, _ := c.GetFloat("payout")
	postbackRate, _ := c.GetInt("rate")
	offerID, _ := c.GetInt("offer_id")
	campID, _ := c.GetInt("camp_id")
	fmt.Println(affName, promoter, postbackURL, offerID,campID)
	postback := new(dimoco.Postback)
	if offerID != 0 && campID != 0 {
		err := postback.CheckOfferID(int64(offerID))
		if err == nil && postback.ID != 0 {
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
				c.Ctx.WriteString("ERROR,插入postbak失败")
				c.StopRun()
			}
		}
	} else {
		c.Ctx.WriteString("ERROR,offerID 是空")
		c.StopRun()
	}
	c.Ctx.WriteString("SUCCESS")
}
