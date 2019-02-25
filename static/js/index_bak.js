
var QualityDatas = {};
var viewDatas = {};
var pageDatas = {};
var pageViewDatas = {};
var data_aff = {};
var subMoData = {};
var  every_date = {
	"lists":{
		"date":"日期",
		"转化":[{"kimia":["1_click"]},{"kimia2":["2_click"]},{"kimia3":["1_click"]},{"kimia4":["1_click","2_click"]},{"kimia5":["1_click","2_click"]},{"kimia6":["1_click"]},{"kimia7":["2_click"]},{"kimia8":["1_click","2_click"]},{"kimia9":["1_click","2_click"]},{"kimia10":["1_click","2_click"]},{"kimia11":["1_click","2_click"]},{"totalSub":["total"]}],
		"花费":[{"kimia":["1_click"]},{"kimia2":["2_click"]},{"kimia3":["1_click"]},{"kimia4":["1_click","2_click"]},{"kimia5":["1_click","2_click"]},{"kimia6":["1_click"]},{"kimia7":["2_click"]},{"kimia8":["1_click","2_click"]},{"kimia9":["1_click","2_click"]},{"kimia10":["1_click","2_click"]},{"kimia11":["1_click","2_click"]},{"totalPostbak":["total"]}],
		"退订":[{"kimia":["1_click"]},{"kimia2":["2_click"]},{"kimia3":["1_click"]},{"kimia4":["1_click","2_click"]},{"kimia5":["1_click","2_click"]},{"kimia6":["1_click"]},{"kimia7":["2_click"]},{"kimia8":["1_click","2_click"]},{"kimia9":["1_click","2_click"]},{"kimia10":["1_click","2_click"]},{"kimia11":["1_click","2_click"]},{"totalUnSub":["total"]}],
		"留存":"",
		"扣费成功数":"",
		"扣费成功率":"",
		"扣费金额":"美金",
		"Orange收入":"3.466/4.5",
		"Three 收入":"3.318/4.5",
		"T-mobile收入":"2.776/4.5",
		"Virgin收入":"1.903/4.5",
		"总收入":"分成比例",
		"盈亏":""
	},
	"data":[

		{
			"date":"2017-12-12",
			"subData":[10210,204010,10210,204010,10210,204010,10210,204010,10210,204010,10210,204010],
			"Postback":[1020,202010,10210,204010,10210,204010,10210,204010,10210,204010,10210,204010],
			"unsubData":[1020,202010,10210,204010,10210,204010,10210,204010,10210,204010,10210,204010],
			"留存":"124321",
			"扣费成功数":"12341235",
			"扣费成功率":"1325132",
			"扣费金额":"5321",
			"Orange收入":"43124",
			"Three收入":"353415",
			"Tmobile收入":"5315",
			"Virgin收入":"1.9031234",
			"总收入":"13412356",
			"盈亏":"134134"
		},
		{
			"date":"2017-12-13",
			"subData":[10120,200102,10210,204010,10210,204010,10210,204010,10210,204010,10210,204010],
			"Postback":[10220,200140,10210,204010,10210,204010,10210,204010,10210,204010,10210,204010],
			"unsubData":[10420,220010,10210,204010,10210,204010,10210,204010,10210,204010,10210,204010],
			"留存":"53213421",
			"扣费成功数":"531243",
			"扣费成功率":"5123",
			"扣费金额":"13564",
			"Orange收入":"33215",
			"Three收入":"532145",
			"Tmobile收入":"2135.5",
			"Virgin收入":"15123.9123",
			"总收入":"51234",
			"盈亏":"512351"
		}
		]
}


function getAffiliateData() {
	$.ajax({
		url: 'http://127.0.0.1:8085/aff_data',
		type: 'GET',
		data: data_aff,
		dataType: "json",
		success: function (result) {
			var data = result.data;
			console.log(data);
			var aff_html = [];
			$.each(data, function (i, c) {
				var pp = [];
				for (var t = 0; t < c.Aff_data.length; t++) {
					for (var a = 0; a < c.Aff_data[t].Ser_list.length; a++) {
						pp.push(c.Aff_data[t].Ser_list[a].Servername);
					}
				}
				console.log(pp);
				aff_html.push('<tr><td rowspan="' + pp.length + '">' + c.AffName + '</td>');
				$.each(c.Aff_data, function (n, q) {
					console.log(q.Ser_list.length + "+b-length");
					aff_html.push('<td rowspan="' + q.Ser_list.length + '">' + q.Pubname + '</td>');
					$.each(q.Ser_list, function (o, p) {
						var activeuser = p.Total_num - p.Unsub_num;
						aff_html.push('<td>' + p.Servername + '</td><td>' + p.Click_num + '</td><td>' + p.Total_num + '</td><td>' + p.PostNum + '</td><td>' + activeuser + '</td><td>' + p.Unsub_num + '</td><td>' + p.SuccessMT_Num + '</td><td>' + p.FailtMT_Num + '</td><td>' + p.Churn_rate + '</td></tr>');
					});
				});
			});
			$("#aff_content").html(aff_html.join(""));
			$("#daochu2").html("<a id='down2' class='btn btn-search' onclick=\"downloadTable2Excal(data_aff.aff_name,data_aff.serverType,data_aff.operator,data_aff.start_time,data_aff.end_time)\">导出表格</a>")
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	});

}

function getMoDateData() {
	$.ajax({
		url: 'http://127.0.0.1:8085/sub/mo_data',
		type: 'GET',
		data: subMoData,
		dataType: "json",
		success: function (result) {
			var data = result.data;
			console.log(data);
			var aff_html = [];
			$.each(data, function (i, c) {
				var pp = [];
				for (var t = 0; t < c.ServiceList.length; t++) {
					for (var a = 0; a < c.ServiceList[t].AffSubData.length; a++) {
						pp.push(c.ServiceList[t].AffSubData[a].AffName);
					}
				}
				console.log(pp);
				aff_html.push('<tr><td rowspan="' + pp.length + '">' + c.Operator + '</td>');
				$.each(c.ServiceList, function (n, q) {
					console.log(q.AffSubData.length + "+b-length");
					aff_html.push('<td rowspan="' + q.AffSubData.length + '">' + q.ServiceName + '</td><td rowspan="' + q.AffSubData.length + '">' + q.Price+ '</td>');
					$.each(q.AffSubData, function (o, p) {
						// var activeuser = p.Total_num - p.Unsub_num;
						aff_html.push('<td>' + p.AffName + '</td><td>' + p.SubNum+ '</td><td>' + p.UnsubNum + '</td><td>' + p.PostbackNum + '</td><td>' + p.FailedMt + '</td><td>' + p.SuccessMt + '</td><td>' + p.Amount + '</td></tr>');
					});
				});
			});
			$("#sub_aff_data").html(aff_html.join(""));

			$("#daochu4").html("<a id='down4' class='btn btn-search' onclick=\"downloadTable4Excal(subMoData.aff_name,subMoData.serverType,subMoData.operator,subMoData.start_sub,subMoData.end_sub,subMoData.start_date,subMoData.end_date)\">导出表格</a>")
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	});

}


function getQualitypage() {
	$.ajax({
		url: 'http://127.0.0.1:8085/world_play/quality',
		type: 'GET',
		data: QualityDatas,
		dataType: "json",
		success: function (result) {
			var data = result;
			console.log(data);
			//table 2
			var aff_html = [];
			$.each(data.data, function (i, c) {
				aff_html.push('<tr><td>' + c.Date + '</td><td>' + c.TotalSubNum + '</td><td>' + c.PostbackNum + '</td><td>' + c.UnsubNum + '</td><td>' + c.ActivateNum + '</td><td>' + c.TotalMt + '</td><td>' + c.RenewNum + '</td><td>' + c.MtFailed + '</td></tr>');
			});
			$("#qm_content").html(aff_html.join(""));
			$("#daochu3").html("<a id='down3' class='btn btn-search' onclick=\"downloadTable3Excal(QualityDatas.aff_name,QualityDatas.pub_id,QualityDatas.serverType,QualityDatas.operator,QualityDatas.sub_date,QualityDatas.end_date)\">导出表格</a>")
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	})
};


function GetEveryDaySubPage() {
	var list_title1 =[];
	var list_title2 =[];
	var list_title3 =[];
	$.ajax({
		url: 'http://127.0.0.1:8085/sub/everyday/data?date=2017-12',
		type: 'GET',
		data: QualityDatas,
		dataType: "json",
		success: function (result) {
			var data = result;
			console.log(data);
			//table 2
			var aff_html = [];
			$.each(data.data, function (i, c) {
				aff_html.push('<tr><td>' + c.Date + '</td><td>' + c.TotalSubNum + '</td><td>' + c.PostbackNum + '</td><td>' + c.UnsubNum + '</td><td>' + c.ActivateNum + '</td><td>' + c.TotalMt + '</td><td>' + c.RenewNum + '</td><td>' + c.MtFailed + '</td></tr>');
			});
			$("#qm_content").html(aff_html.join(""));
			$("#daochu3").html("<a id='down3' class='btn btn-search' onclick=\"downloadTable3Excal(QualityDatas.aff_name,QualityDatas.pub_id,QualityDatas.serverType,QualityDatas.operator,QualityDatas.sub_date,QualityDatas.end_date)\">导出表格</a>")
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	})


	for (var key in every_date.lists){
		if (key === "转化" || key === "花费" || key === "退订"){
			list_title1.push('<th colspan='+every_date.lists[key].length+'>'+ key + '</th>');
			$.each(every_date.lists[key],function (i,c) {
				list_title2.push('<th colspan="'+c[Object.keys(c)].length+'">' + Object.keys(c)+ '</th>')
				$.each(c[Object.keys(c)],function (i,key1) {
					list_title3.push('<th>'+key1+'</th>')
                })

       		 })
		}else{
			list_title1.push('<th>'+ key + '</th>');
			list_title2.push('<th>' + every_date.lists[key]+ '</th>')
			list_title3.push('<th></th>')
		}
	}


	$("#titleName").html('<tr>' + list_title1.join("") +'</tr>' + '<tr>' + list_title2.join("") +'</tr>'+'<tr>' + list_title3.join("") +'</tr>');

	var lists2 = [];


	$.each(every_date.data,function (i,c) {
		lists2.push('<tr><td>' + c.date + '</td>');
		$.each(c.subData,function (i,c2) {
			lists2.push('<td>' + c2 + '</td>')
        });

		lists2.push('<td style="background-color:#bedead">' + 1111 + '</td>')
		$.each(c.Postback,function (i,c2) {
			lists2.push('<td>' + c2 + '</td>')
        });
		lists2.push('<td style="background-color:#bedead">' + 1122211 + '</td>')
		$.each(c.unsubData,function (i,c2) {
			lists2.push('<td >' + c2 + '</td>')
        });
		lists2.push('<td style="background-color:#bedead">' + 11133231 + '</td>')
		lists2.push('<td>' + c.留存 + '</td>')
		lists2.push('<td>' + c.扣费成功数 + '</td>')
		lists2.push('<td>' + c.扣费成功率 + '</td>')
		lists2.push('<td>' + c.扣费金额 + '</td>')
		lists2.push('<td>' + c.Orange收入 + '</td>')
		lists2.push('<td>' + c.Three收入 + '</td>')
		lists2.push('<td>' + c.Tmobile收入 + '</td>')
		lists2.push('<td>' + c.Virgin收入 + '</td>')
		lists2.push('<td>' + c.总收入 + '</td>')
		lists2.push('<td>' + c.盈亏 + '</td></tr>')
    })
	$("#subtotal").html(lists2.join(""));
}


$(document).ready(function () {
	$(function () {
		$(".datepicker").datetimepicker({
			showSecond: true,
			showMillisec: true,
			dateFormat: "yy-mm-dd"
		});
	});
	$("#Subscriber").show();
	$("#Quality").hide();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#EveryDaySubQuality").hide();
});
$("#searchThird").click(function () {
	var start_time = document.getElementById("start_time_aff").value ? document.getElementById("start_time_aff").value : GetSevenDayDate();
	var end_time = document.getElementById("end_time_aff").value ? document.getElementById("end_time_aff").value : NowDate();

	data_aff = {
		"start_time": start_time,
		"end_time": end_time,
		"operator": $("#Opearator option:selected").text(),
		"aff_name": $("#Aff_Name option:selected").text(),
		"serverType": $("#service_type option:selected").text(),
		"clickType": $("#click_type option:selected").text()
	};
	getAffiliateData();
});



$("#searchMoAffDate").click(function () {
	var start_sub_date = document.getElementById("start_sub_date").value ? document.getElementById("start_sub_date").value : GetSevenDayDate();
	var end_sub_date = document.getElementById("end_sub_date").value ? document.getElementById("end_sub_date").value : NowDate();
	var start_date = document.getElementById("start_date").value ? document.getElementById("start_date").value : GetSevenDayDate();
	var end_date = document.getElementById("end_date").value ? document.getElementById("end_date").value : NowDate();

	subMoData = {
		"start_sub": start_sub_date,
		"end_sub": end_sub_date,
		"start_date":start_date,
		"end_date":end_date,
		"operator": $("#Opearator2 option:selected").text(),
		"aff_name": $("#Aff_Name2 option:selected").text(),
		"service_type": $("#service_type2 option:selected").text(),
		"clickType": $("#click_type2 option:selected").text(),
	};
	getMoDateData();
});



//change nav page
$("#view_title").click(function () {
	$("#Quality").hide();
	$("#Subscriber").show();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#EveryDaySubQuality").hide();

	$("#back3").removeClass("background");
	$("#back4").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").removeClass("background");
	$("#back1").addClass("background");

});
$("#qm_title").click(function () {
	$("#Subscriber").hide();
	$("#Quality").show();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#EveryDaySubQuality").hide();

	$("#back1").removeClass("background");
	$("#back4").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").removeClass("background");
	$("#back3").addClass("background");

	QualityDatas = {
		"aff_name": "All",
		"operator": "All",
		"date": "2017-09-04"
	};
	getQualitypage();
});
$("#aff_title").click(function () {
	$("#Subscriber").hide();
	$("#Quality").hide();
	$("#Affiliate").show();
	$("#AffSubQuality").hide();
	$("#Everydaysubquality").hide();

	$("#back3").removeClass("background");
	$("#back4").removeClass("background");
	$("#back1").removeClass("background");
	$("#back5").removeClass("background");
	$("#back2").addClass("background");

	data_aff = {
		"start_time": GetSevenDayDate(),
		"end_time": NowDate(),
		"operator": "All",
		"aff_name": "All",
		"serverType": "All"
	};
	getAffiliateData();
});

$("#sub_qm_title").click(function () {
	$("#Subscriber").hide();
	$("#Quality").hide();
	$("#Affiliate").hide();
	$("#EveryDaySubQuality").hide();
	$("#AffSubQuality").show();


	$("#back3").removeClass("background");
	$("#back1").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").removeClass("background");
	$("#back4").addClass("background");




	var start_sub_date = document.getElementById("start_sub_date").value ? document.getElementById("start_sub_date").value : GetSevenDayDate();
	var end_sub_date = document.getElementById("end_sub_date").value ? document.getElementById("end_sub_date").value : NowDate();
	var start_date = document.getElementById("start_date").value ? document.getElementById("start_date").value : GetSevenDayDate();
	var end_date = document.getElementById("end_date").value ? document.getElementById("end_date").value : NowDate();
	subMoData = {
		"start_sub": start_sub_date,
		"end_sub": end_sub_date,
		"start_date":start_date,
		"end_date":end_date,
		"operator": $("#Opearator2 option:selected").text(),
		"aff_name": $("#Aff_Name2 option:selected").text(),
		"service_type": $("#service_type2 option:selected").text()
	};
	getMoDateData();
});



$("#every_sub_qm").click(function () {
	$("#Subscriber").hide();
	$("#Quality").hide();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#EveryDaySubQuality").show();


	$("#back3").removeClass("background");
	$("#back4").removeClass("background");
	$("#back1").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").addClass("background");

	GetEveryDaySubPage()
	// getAffiliateData();
});


// view page search
$("#query12").click(function () {
	viewDatas = {
		"operator": $("#telco_view option:selected").text(),
		"start_time": document.getElementById("start_time_view").value ? document.getElementById("start_time_view").value : GetSevenDayDate(),
		"end_time": document.getElementById("end_time_view").value ? document.getElementById("end_time_view").value : NowDate(),
		"aff_name": $("#affName").val(),
		"service_type": $("#serviceType").val(),
		"pubid": $("#pubId option:selected").text(),
		"clickType":$("#clickType option:selected").text()
	};

	$.ajax({
		url: 'http://127.0.0.1:8085/aff_mt',
		type: 'GET',
		data:viewDatas,
		dataType: "json",
		success: function (result) {
			var data = result;
			var totalMt = data.datas.SuccessMt + data.datas.FailedMt;
			$("#viewtb1").html("<tr><td>" + data.datas.TotalSub + "</td><td>" + data.datas.TotalUnsub + "</td><td>" + data.datas.TotalPostback + "</td><td>" + totalMt + "</td><td>" + data.datas.SuccessMt + "</td><td>" + data.datas.FailedMt + "</td><td>" + data.datas.MtRate + "</td></tr>");
			$("#daochu").html("<a id='down1' class='btn btn-search' onclick=\"downloadTable1Excal(viewDatas.aff_name,viewDatas.pubid,viewDatas.service_type,viewDatas.operator,viewDatas.start_time,viewDatas.end_time)\">导出表格</a>")
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	})
});

// Quality page search


$("#searchQm").click(function () {
	event.preventDefault();
	var time = document.getElementById("Quality_start_time").value;
	time = time.substr(0, 10);
	var end_date = document.getElementById("Quality_end_time").value;
	time = time.substr(0, 10);
	var time = time ? time : GetSevenDayDate();
	var end_date = end_date ? end_date : NowDate();
	QualityDatas = {
		"aff_name": $("#Quality_Aff_Name option:selected").text(),
		"operator": $("#Quality_Opearator option:selected").text(),
		"sub_date": time,
		"end_date": end_date,
		"serverType": $("#Quality_service_type option:selected").text(),
		"pub_id": $("#pud_select option:selected").text(),
		"clickType": $("#Quality_click_type option:selected").text()
	};
	console.log(QualityDatas);
	getQualitypage();
});
// traffic page clear
$("#clearData").click(function () {
	event.preventDefault();
	location.reload();
});

$("#next").click(function () {
	event.preventDefault();
	turnTrafficPage();
});
$("#nav").click(function () {
	event.preventDefault();
	turnTrafficPage();
});
$("#go").click(function () {
	event.preventDefault();
	turnTrafficPage();
});
$("#prev").click(function () {
	event.preventDefault();
	turnTrafficPage();
});
$("#first").click(function () {
	event.preventDefault();
	turnTrafficPage();
});

$("#next1").click(function () {
	event.preventDefault();
	turnViewPage();
});
$("#nav1").click(function () {
	event.preventDefault();
	turnViewPage();
});
$("#go1").click(function () {
	event.preventDefault();
	turnViewPage();
});
$("#prev1").click(function () {
	event.preventDefault();
	turnViewPage();
});
$("#first1").click(function () {
	event.preventDefault();
	turnViewPage();
});


function GetSevenDayDate() {   //  获取7天前日期
	var myDate = new Date(); //获取今天日期
	var dateTemp;
	myDate.setDate(myDate.getDate() - 7);
	dateTemp = myDate.getFullYear() + "-" + (myDate.getMonth() + 1) + "-" + myDate.getDate();
	return dateTemp
}


function NowDate() {
	var myDate = new Date(); //获取今天日期
	var dateTemp;
	dateTemp = myDate.getFullYear() + "-" + (myDate.getMonth() + 1) + "-" + myDate.getDate();
	return dateTemp
}

function changeMe1() {
	$.ajax({
		url: 'http://127.0.0.1:8085/get_pubid?aff_name=' + $("#affName").val(),
		type: 'GET',
		dataType: "json",
		success: function (result) {
			var data = result.data;
			console.log(data);
			var aff_html = [];
			aff_html.push('<select class="selectpicker" id="pud_select" data-live-search="true" title="Please select a lunch ..."><option>All</option>')
			$.each(data, function (o, p) {
				aff_html.push('<option>' + p + '</option>');
			});
			aff_html.push('</select>')
			console.log(aff_html)
			$("#pubId").html(aff_html.join(""));
			$('.selectpicker').selectpicker({
                'selectedText': 'cat'
            });
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	});
}

function changeMe() {
	var a = $("#Quality_Aff_Name").val();
	$.ajax({
		url: 'http://127.0.0.1:8085/get_pubid?aff_name=' + a,
		type: 'GET',
		dataType: "json",
		success: function (result) {
			var data = result.data;
			console.log(data);
			var aff_html = [];
			aff_html.push('<select class="selectpicker" id="pud_select" data-live-search="true" title="Please select a lunch ..."><option>All</option>')
			$.each(data, function (o, p) {
				aff_html.push('<option>' + p + '</option>');
			});
			aff_html.push('</select>')
			console.log(aff_html)
			$("#ziqudao").html(aff_html.join(""));
			$('.selectpicker').selectpicker({
                'selectedText': 'cat'
            });
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	});
}



$('.form_datetime2').datetimepicker({
    minView: "month", //选择日期后，不会再跳转去选择时分秒
    language:  'zh-CN',
    format: 'yyyy-mm-dd',
    todayBtn:  1,
    autoclose: 1,
});



function downloadTable1Excal(aff_name,PubId,service_type,operator,start_sub, end_sub) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>AffName:"+aff_name+"  |PubId:"+PubId+"   |ServiceType:"+service_type+"  |Operator:"+operator+"   |Start Sub Date: "+start_sub+" |End Sub Date: "+end_sub+"</h4>" + document.getElementById("table1").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文件内容的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down1");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "扣费数量.xls";
}

function downloadTable2Excal(aff_name,service_type,operator,start_sub, end_sub) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>AffName:"+aff_name+"   |ServiceType:"+service_type+"  |Operator:"+operator+"   |Start Sub Date: "+start_sub+" |End Sub Date: "+end_sub+"</h4>" + document.getElementById("table2").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文件内容的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down2");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "渠道质量.xls";
}


function downloadTable3Excal(aff_name,PubId,service_type,operator,start_sub, end_sub) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>AffName:"+aff_name+"   |PubId:"+PubId+"      |ServiceType:"+service_type+"  |Operator:"+operator+"   |Start Sub Date: "+start_sub+" |End Sub Date: "+end_sub+"</h4>" + document.getElementById("table3").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文件内容的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down3");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "持续扣费情况.xls";
}


function downloadTable4Excal(aff_name,service_type,operator,start_sub, end_sub,start,end) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>AffName:"+aff_name+"     |ServiceType:"+service_type+"  |Operator:"+operator+"   |Start Sub Date: "+start_sub+" |End Sub Date: "+end_sub+"   |Search Start Date: "+start+"   |Search End Date: "+end+"</h4>" +document.getElementById("table4").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down4");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "订阅留存.xls"
}