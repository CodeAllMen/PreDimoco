
var QualityDatas = {};
var viewDatas = {};
var pageDatas = {};
var pageViewDatas = {};
var data_aff = {};
var subMoData = {};

function getAffiliateData() {
	$.ajax({
		url: 'http://cpx2.allcpx.com/aff_data',
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
		url: 'http://cpx2.allcpx.com/sub/mo_data',
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
		url: 'http://cpx2.allcpx.com/world_play/quality',
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

});
$("#searchThird").click(function () {
	var start_time = document.getElementById("start_time_aff").value ? document.getElementById("start_time_aff").value : GetSevenDayDate();
	var end_time = document.getElementById("end_time_aff").value ? document.getElementById("end_time_aff").value : NowDate();

	data_aff = {
		"start_time": start_time,
		"end_time": end_time,
		"operator": $("#Opearator option:selected").text(),
		"aff_name": $("#Aff_Name option:selected").text(),
		"serverType": $("#service_type option:selected").text()
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
		"service_type": $("#service_type2 option:selected").text()
	};
	getMoDateData();
});


//change nav page
$("#view_title").click(function () {
	$("#Quality").hide();
	$("#Subscriber").show();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#everydaysubquality").hide();

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
	$("#everydaysubquality").hide();

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
	$("#everydaysubquality").hide();

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
	$("#AffSubQuality").show();
	$("#everydaysubquality").hide();

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
	$("#everydaysubquality").show();


	$("#back3").removeClass("background");
	$("#back4").removeClass("background");
	$("#back1").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").addClass("background");
	data_aff = {
		"start_time": GetSevenDayDate(),
		"end_time": NowDate(),
		"operator": "All",
		"aff_name": "All",
		"serverType": "All"
	};
	getAffiliateData();
});





// view page search
$("#query12").click(function () {
	viewDatas = {
		"operator": $("#telco_view option:selected").text(),
		"start_time": document.getElementById("start_time_view").value ? document.getElementById("start_time_view").value : GetSevenDayDate(),
		"end_time": document.getElementById("end_time_view").value ? document.getElementById("end_time_view").value : NowDate(),
		"aff_name": $("#affName").val(),
		"service_type": $("#serviceType").val(),
		"pubid": $("#pubId option:selected").text()
	};

	$.ajax({
		url: 'http://cpx2.allcpx.com/aff_mt',
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

function pubIdShow() {
	var a = $("#affName").val();
	$.ajax({
		url: 'http://cpx2.allcpx.com/get_pubid?aff_name=' + a,
		type: 'GET',
		dataType: "json",
		success: function (result) {
			var data = result.data;
			console.log(data);
			var aff_html = [];
			aff_html.push('<label class="checkbox-inline"><input type="checkbox" onclick="checkAllOne()"> 全选</label>')
			$.each(data, function (o, p) {
				aff_html.push('<label class="checkbox-inline"><input type="checkbox"  name="check_q" value="' + p + '">' + p + '</label>');
			});
			$("#pubId").html(aff_html.join(""));

		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	});
}

function checkAllOne() {
	$("[name='check_q']").attr("checked", 'true');
}


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
		url: 'http://cpx2.allcpx.com/get_pubid?aff_name=' + $("#affName").val(),
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
		url: 'http://cpx2.allcpx.com/get_pubid?aff_name=' + a,
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