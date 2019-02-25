
var QualityDatas = {};
var viewDatas = {};
var pageDatas = {};
var pageViewDatas = {};
var data_aff ={};

function getAffiliateData() {
	var start_time = document.getElementById("start_time_aff").value ? document.getElementById("start_time_aff").value : GetSevenDayDate();
	var end_time = document.getElementById("end_time_aff").value ? document.getElementById("end_time_aff").value : NowDate();

	data_aff = {
		"start_time": start_time,
		"end_time": end_time,
		"operator": $("#Opearator option:selected").text(),
		"aff_name": $("#Aff_Name option:selected").text(),
		"serverType": $("#service_type option:selected").text()
	};
	$.ajax({
		url: 'http://allcpx.com/aff_data',
		type: 'GET',
		data: data_aff,
		dataType: "json",
		success: function(result){
			var data = result.data;
			console.log(data);
			var aff_html = [];
			$.each(data, function(i,c) {
				var pp = [];
				for (var t = 0; t<c.Aff_data.length; t++) {
					for (var a = 0; a< c.Aff_data[t].Ser_list.length; a++) {
						pp.push(c.Aff_data[t].Ser_list[a].Servername);
					}
				}
				console.log(pp);
				// console.log(c.Aff_data.length + "rhgiuqhgui");
				aff_html.push('<tr><td rowspan="'+pp.length+'">'+c.Name+'</td>');
					$.each(c.Aff_data, function(n, q) {
						console.log(q.Ser_list.length+"+b-length");
						aff_html.push('<td rowspan="'+q.Ser_list.length+'">'+ q.Pubname +'</td>');
						$.each(q.Ser_list, function(o, p) {
							aff_html.push('<td>'+ p.Servername+'</td><td>'+ p.Click_num +'</td><td>'+p.Total_num+'</td><td>'+p.PostNum+'</td><td>'+ p.Active_num+'</td><td>'+ p.Unsub_num +'</td><td>'+ p.SuccessMT_Num+'</td><td>'+ p.FailtMT_Num +'</td><td>'+ p.Churn_rate  +'</td></tr>');
						});
					});
			});
			$("#aff_content").html(aff_html.join(""));
		}.bind(this),
		error: function(error) {
			console.log(error);
		}
	});

}

function getQualitypage() {
	$.ajax({
		url: 'http://allcpx.com/world_play/quality',
		type: 'GET',
		data: QualityDatas,
		dataType: "json",
		success: function(result){
			var data = result;
			console.log(data);
			//table 2 
			var aff_html = [];
			$.each(data.data, function(i, c) {
				aff_html.push('<tr><td>'+ c.Date +'</td><td>'+ c.TotalSubNum +'</td><td>'+ c.PostbackNum +'</td><td>'+ c.UnsubNum +'</td><td>'+ c.ActivateNum +'</td><td>'+ c.TotalMt +'</td><td>'+ c.RenewNum +'</td><td>'+ c.MtFailed +'</td></tr>');
			});
			$("#qm_content").html(aff_html.join(""));
		}.bind(this),
		error: function(error) {
			console.log(error);
		}
	})	
};

$(document).ready(function(){
	$( function() {
	    $( ".datepicker" ).datetimepicker({
	    	showSecond: true,
			showMillisec: true,
			// timeFormat: "HH:mm:ss",
            dateFormat: "yy-mm-dd"
	    });
	});

	$("#Quality").hide();
	$("#Affiliate").hide();

});
$("#searchThird").click(function(){
	getAffiliateData();
})
//change nav page
$("#view_title").click(function(){
	$("#Quality").hide();
	$("#Affiliate").hide();
});
$("#qm_title").click(function(){
	$("#Quality").show();
	$("#Affiliate").hide();
	QualityDatas = {
		"aff_name": "All",
		"operator": "All",
		"date": "2017-09-04"
	};
	getQualitypage();
});
$("#aff_title").click(function(){
	$("#Quality").hide();
	$("#Affiliate").show();
	getAffiliateData();
});
// view page search

// Quality page search


$("#searchQm").click(function(){
	event.preventDefault();
	var time = document.getElementById("Quality_start_time").value;
	time = time.substr(0,10);
	var end_date = document.getElementById("Quality_end_time").value;
	time = time.substr(0,10);
	var time = time ? time : GetSevenDayDate();
	var end_date = end_date ? end_date : NowDate();
	QualityDatas = {
		"aff_name": $("#Quality_Aff_Name option:selected").text(),
		"operator": $("#Quality_Opearator option:selected").text(),
		"sub_date": time,
		"end_date": end_date,
		"serverType": $("#Quality_service_type option:selected").text()
	};
	console.log(QualityDatas);
	getQualitypage();
});
// traffic page clear
$("#clearData").click(function(){
	event.preventDefault();
	location.reload();	
});

$("#next").click(function(){
	event.preventDefault();	
	turnTrafficPage();
});
$("#nav").click(function(){
	event.preventDefault();	
	turnTrafficPage();
});
$("#go").click(function(){
	event.preventDefault();	
	turnTrafficPage();
});
$("#prev").click(function(){
	event.preventDefault();	
	turnTrafficPage();
});
$("#first").click(function(){
	event.preventDefault();	
	turnTrafficPage();
});

$("#next1").click(function(){
	event.preventDefault();	
	turnViewPage();
});
$("#nav1").click(function(){
	event.preventDefault();	
	turnViewPage();
});
$("#go1").click(function(){
	event.preventDefault();	
	turnViewPage();
});
$("#prev1").click(function(){
	event.preventDefault();	
	turnViewPage();
});
$("#first1").click(function(){
	event.preventDefault();	
	turnViewPage();
});



function GetSevenDayDate() {
	var myDate = new Date(); //获取今天日期
	var dateTemp;
	myDate.setDate(myDate.getDate() - 7);
	dateTemp = myDate.getFullYear()+"-"+ (myDate.getMonth()+1)+"-"+myDate.getDate();
	return dateTemp
}


function NowDate() {
	var myDate = new Date(); //获取今天日期
	var dateTemp;
	// alert(length(toString(myDate.getMonth()+1)));
	dateTemp = myDate.getFullYear()+"-"+ (myDate.getMonth()+1)+"-"+myDate.getDate();
	return dateTemp
}