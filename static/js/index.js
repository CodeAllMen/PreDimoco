var QualityDatas = {};
var viewDatas = {};
var pageDatas = {};
var subeveryDayData = {};
var data_aff = {};
var subMoData = {};
var every_date = {
	"columnData": {
		// "date": "日期",
		"转化": "",
		"postback回传数": "",
		"花费": "",
		"MT成功数": "",
		"退订": "",
		"留存": "",
		"扣费成功数": "",
		"扣费成功率": "",
		"扣费金额": "美金",
		"Tele2收入": "2.773/4.5",
		"Vodafone 收入": "2.622/4.5",
		"T-mobile收入": "2.776/4.5",
		"KPN收入": "2.834/4.5",
		"Telfort收入": "2.879/4.5",
		"当天收入": "分成比例",
		"当天": ["转化", "扣费成功数", "退订", "花费", "收入"],
		"累计": ["转化", "留存", "扣费成功数", "扣费成功率", "花费", "收入", "盈亏"],
		"上月累计收入": "XXxx",
		"上月累计花费": "xxxxx",
		"上月盈亏": "xxxxxx",
	}
}
var firstNum = false
var firstNum2 = false
var firstNum3 = false
var firstNum4 = false
var firstNum5 = false
var char_data = []
var char_data1 = []

function getAffiliateData() {
	$.ajax({
		url: 'http://pl.leadernethksp.com/aff_data',
		type: 'GET',
		data: data_aff,
		dataType: "json",
		success: function (result) {
			var data = result.data;
			var aff_html = [];
			$.each(data, function (i, c) {
				var pp = [];
				for (var t = 0; t < c.Aff_data.length; t++) {
					for (var a = 0; a < c.Aff_data[t].Ser_list.length; a++) {
						pp.push(c.Aff_data[t].Ser_list[a].Servername);
					}
				}
				aff_html.push('<tr><td rowspan="' + pp.length + '">' + c.AffName + '</td>');
				$.each(c.Aff_data, function (n, q) {
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


function getQualitypage() {
	$.ajax({
		url: 'http://pl.leadernethksp.com/quality',
		type: 'GET',
		data: QualityDatas,
		dataType: "json",
		success: function (result) {
			var data = result;
			var aff_html = [];
			$.each(data.tableData, function (i, c) {
				var color = ""
				if (i === 0){
					color = "background-color:#FF3333"
				}
				aff_html.push('<tr style="'+color+'"><td>' + c.Date + '</td><td>' + c.TotalSubNum + '</td><td>' + c.PostbackNum + '</td><td>' + c.UnsubNum + '</td><td>' + c.ActivateNum + '</td><td>' + c.TotalMt + '</td><td>' + c.RenewNum + '</td><td>' + c.MtFailed + '</td><td>' + c.DayRevenue + '</td><td>' + c.Spend + '</td><td>' + c.Revenue + '</td><td>' + c.ProfitAndLoss + '</td></tr>');
			});
			$("#qm_content").html(aff_html.join(""));
			char_data = data.chartData
			ShowChart2(data.chartData)
			ShowChart3(data.chartData)

			$("#daochu3").html("<a id='down3' class='btn btn-search' onclick=\"downloadTable3Excal(QualityDatas.aff_name,QualityDatas.pub_id,QualityDatas.serverType,QualityDatas.operator,QualityDatas.sub_date,QualityDatas.end_date)\">导出表格</a>")
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	})
};


function getEerytimeQualitypage() {
	$.ajax({
		url: 'http://pl.leadernethksp.com/sub/mo_data',
		type: 'GET',
		data: subMoData,
		dataType: "json",
		success: function (result) {
			var data = result;
			var aff_html = [];
			$.each(data.tableData, function (i, c) {
				var color = ""
				if (i === 0){
					color = "background-color:#FF3333"
				}
				aff_html.push('<tr style="'+color+'"><td>' + c.Date + '</td><td>' + c.TotalSubNum + '</td><td>' + c.PostbackNum + '</td><td>' + c.UnsubNum + '</td><td>' + c.ActivateNum + '</td><td>' + c.TotalMt + '</td><td>' + c.RenewNum + '</td><td>' + c.MtFailed + '</td><td>' + c.DaySpend + '</td><td>' + c.DayRevenue + '</td><td>' + c.Spend + '</td><td>' + c.Revenue + '</td><td>' + c.ProfitAndLoss + '</td></tr>');
			});
			$("#qm_content1").html(aff_html.join(""));
			char_data1 = data.chartData
			ShowChart4(data.chartData)
			ShowChart5(data.chartData)

			$("#daochu4").html("<a id='down4' class='btn btn-search' onclick=\"downloadTable4Excal(QualityDatas.aff_name,QualityDatas.pub_id,QualityDatas.serverType,QualityDatas.operator,QualityDatas.sub_date,QualityDatas.end_date)\">导出表格</a>")
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	})
};

function ComplaintsData(userMsisdn) {
	var list_title1 = [];
	$.ajax({
		url: 'http://pl.leadernethksp.com/msisdn',
		type: 'GET',
		data: { 'msisdn': userMsisdn },
		dataType: 'json',
		success: function (result) {
			$("#com_msisdn").val(result.data[0].Msisdn)
			$("#com_email").val(result.data[0].Email)
			$("#com_user").val(result.data[0].UserName)
			$("#com_equi").val(result.data[0].EquipmentModel)
			$("#com_aff").val(result.data[0].GuiltyAffName)
			$("#com_pub").val(result.data[0].GuiltyPubid)
			$("#com_des").val(result.data[0].Description)
			$("#Com_level").val(result.data[0].Level)
			$("#Com_time").val(result.data[0].DealWithTime)
			if (result.data[0].Msisdn != "没有此电话号码信息") {
				$("#com_info").removeClass("noexit")
			}
			$.each(result.data, function (i, data) {
				list_title1.push("<tr><td>" + data.Msisdn + "</td><td>" + data.Operator + "</td><td>" + data.SubId + "</td><td>" + data.ClickType + "</td><td>" + data.AffName + "</td><td>" + data.PubId + "</td><td>" + data.ClickId + "</td><td>" + data.PostbackStatus + "</td><td>" + data.MtNum + "</td><td>" + data.Amount + "</td><td>" + data.RackingCode + "</td><td>" + data.Subtime + "</td><td>" + data.Unsubtime + "</td></tr>");
			});
			$("#ComplaintDatas").html(list_title1.join(""))
		}
	})
}


function GetEveryDaySubPage() {
	var list_title1 = [];
	var affNameColumn = [];
	// var clickType =[];
	$.ajax({
		url: 'http://pl.leadernethksp.com/sub/everyday/data',
		type: 'GET',
		data: subeveryDayData,
		dataType: "json",
		success: function (result) {
			var affClickData = result
			for (var key in every_date.columnData) {
				if (key === "转化" || key === "花费" || key === "退订" || key === "postback回传数" || key === "MT成功数") {
					var lens = 1;
					$.each(affClickData.affNameList, function (i, affName) {
						affNameColumn.push('<th><div class="click_num">' + affName + '</div></th>')
						lens += 1
					})
					affNameColumn.push('<th><div class="click_num">合计</div></th>')
					list_title1.push('<th colspan=' + lens + '>' + key + '</th>');
				} else if (key === "当天" || key === "累计") {
					list_title1.push('<th colspan=' + every_date.columnData[key].length + '>' + key + '</th>');
					$.each(every_date.columnData[key], function (i, c) {
						affNameColumn.push('<th><div class="click_num2">' + c + '</div></th>')
					})
				} else {
					if (key === "上月累计收入") {
						affNameColumn.push('<th style="background-color:#00FF99"><div class="click_num2">' + affClickData.lastMonthRevenue + '</div></th>')
					} else if (key === "上月累计花费") {
						affNameColumn.push('<th style="background-color:#00FF99"><div class="click_num2">' + affClickData.lastMouthSpend + '</div></th>')
					} else if (key === "上月盈亏") {
						var ProfitAndLoss = (parseFloat(affClickData.lastMonthRevenue) - parseFloat(affClickData.lastMouthSpend)).toFixed(3)
						affNameColumn.push('<th style="background-color:#00FF99"><div class="click_num2">' + ProfitAndLoss + '</div></th>')
					} else {
						affNameColumn.push('<th><div class="click_num2">' + every_date.columnData[key] + '</div></th>')
					}
					list_title1.push('<th>' + key + '</th>');
				}
			}

			$("#right_table1").html('<tr>' + list_title1.join("") + '</tr>' + '<tr>' + affNameColumn.join("") + '</tr>');

			var lists2 = [];
			var lists2_date = [];

			$.each(affClickData.data, function (i, c) {
				lists2_date.push('<tr><th style="background-color:#efefef">' + c.Date.substr(2, 20) + '</th><tr>');
				lists2.push('<tr>');

				var total_sub = 0     // 转化
				$.each(c.SubData, function (i, subNum) {
					total_sub += parseInt(subNum)
					lists2.push('<th><div class="click_num">' + subNum + '</div></th>')
				});
				lists2.push('<th style="background-color:#bedead"><div class="click_num">' + total_sub + '</div></th>')

				var total_postNum = 0    //  postback回传数
				$.each(c.PostbackData, function (i, postNum) {
					total_postNum += parseInt(postNum)
					lists2.push('<th><div class="click_num">' + postNum + '</div></th>')
				});
				lists2.push('<th style="background-color:#bedead"><div class="click_num">' + total_postNum + '</div></th>')

				var total_postSpend = 0.0     //  花费
				$.each(c.PostbackSpend, function (i, postSpend) {
					total_postSpend += parseFloat(postSpend)
					lists2.push('<th><div class="click_num">' + postSpend + '</div></th>')
				});
				lists2.push('<th style="background-color:#bedead"><div class="click_num">' + total_postSpend + '</div></th>')

				var total_Mt_Num = 0   //  MT成功数
				$.each(c.MtNumData, function (i, mt_num) {
					total_Mt_Num += parseInt(mt_num)
					lists2.push('<th><div class="click_num">' + mt_num + '</div></th>')
				});
				lists2.push('<th style="background-color:#bedead"><div class="click_num">' + total_Mt_Num + '</div></th>')

				var total_Unsub = 0    //  退订
				$.each(c.UnSubData, function (i, unsubNum) {
					total_Unsub += parseInt(unsubNum)
					lists2.push('<th ><div class="click_num">' + unsubNum + '</div></th>')
				});
				lists2.push('<th style="background-color:#bedead"><div class="click_num">' + total_Unsub + '</div></th>')
				lists2.push('<th><div class="click_num2">' + c.Active + '</div></th>')
				lists2.push('<th><div class="click_num2">' + c.SuccessMt + '</div></th>')
				lists2.push('<th><div class="click_num2">' + c.MtRate + '</div></th>')
				lists2.push('<th><div class="click_num2">' + c.Amout + '</div></th>')
				lists2.push('<th><div class="click_num2">' + c.Tele2 + '</div></th>')
				lists2.push('<th><div class="click_num2">' + c.Tmobile + '</div></th>')
				lists2.push('<th><div class="click_num2">' + c.Vodafone + '</div></th>')
				lists2.push('<th><div class="click_num2">' + c.KPN + '</div></th>')
				lists2.push('<th><div class="click_num2">' + c.Telfort + '</div></th>')
				lists2.push('<th><div class="click_num2">' + c.DayRevenue + '</div></th>')
				lists2.push('<th style="background-color:#2BD5D5"><div class="click_num2">' + total_sub + '</div></th>')
				lists2.push('<th style="background-color:#2BD5D5"><div class="click_num2">' + c.SuccessMt + '</div></th>')
				lists2.push('<th style="background-color:#2BD5D5"><div class="click_num2">' + total_Unsub + '</div></th>')
				lists2.push('<th style="background-color:#2BD5D5"><div class="click_num2">' + c.DaySpend + '</div></th>')
		
				lists2.push('<th style="background-color:#2BD5D5"><div class="click_num2">' + c.DayRevenue + '</div></th>')

				lists2.push('<th style="background-color:#5EA287"><div class="click_num2">' + c.GrandTotalSub + '</div></th>')
				lists2.push('<th style="background-color:#5EA287"><div class="click_num2">' + c.Active + '</div></th>')
				lists2.push('<th style="background-color:#5EA287"><div class="click_num2">' + c.TotalSuccessMt + '</div></th>')
				console.log(c.TotalSuccessMt)
				lists2.push('<th style="background-color:#5EA287"><div class="click_num2">' + c.GrandTotalMtRate + '</div></th>')
				lists2.push('<th style="background-color:#5EA287"><div class="click_num2">' + c.GrandTotalSpend + '</div></th>')
				
				lists2.push('<th style="background-color:#5EA287"><div class="click_num2">' + c.GrandTotalRevenue + '</div></th>')
				lists2.push('<th style="background-color:#5EA287"><div class="click_num2">' + c.GrandTotalProfitAndLoss + '</div></th>')
				lists2.push('<th><div class="click_num2"></div></th>')
				lists2.push('<th><div class="click_num2"></div></th>')
				lists2.push('<th><div class="click_num2"></div></th></tr>')
			})
			$("#left_table2").html(lists2_date.join(""));
			$("#right_table2").html(lists2.join(""));
			$("#all_data").removeClass("noexit")

			$("#daochu5").html("<a id='down5' class='btn btn-search' onclick=\"downloadTable5Excal(subeveryDayData.date)\">导出表格</a>")

		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	})
}


function ShowChart() {
	var start_date = document.getElementById("start_time_char").value ? document.getElementById("start_time_char").value : GetSevenDayDate();
	var end_date = document.getElementById("end_time_char").value ? document.getElementById("end_time_char").value : NowDate();
	var date, spend, profitAndLoss, revenue;
	$.ajax({
		url: 'http://pl.leadernethksp.com/day/chart',
		type: 'GET',
		data: {
			'startDate': start_date.substr(0, 10),
			"endDate": end_date.substr(0, 10),
			"operator": $("#Opearator_char").val(),
			"serverType": $("#Service_type_char").val(),
			"aff_name": $("#Aff_Name_char").val(),
			"clickType": $("#Click_type_char").val(),
			"pub_id": $("#Pud_select_char option:selected").text(),
		},
		dataType: 'json',
		success: function (result) {
			date = result.date
			spend = result.spend
			profitAndLoss = result.profitAndLoss
			revenue = result.revenue
			var optionInit = {
				color: ['#FF0000', '#87cefa', '#33CC52', '#32cd32', '#6495ed'],
				backgroundColor: '#ffffff',
				title: {
					text: 'txtnation每日盈亏折线图',
					subtext: ''
				},
				tooltip: {
					trigger: 'axis'
				},
				legend: {
					data: ['花费', '收入', '盈利']
				},
				toolbox: {
					show: true,
					feature: {
						dataView: { show: true, readOnly: true },
						// magicType: { show: true, type: ['line', 'bar', 'stack', 'tiled'] },
						restore: { show: true },
						saveAsImage: { show: true }
					}
				},
				calculable: true,
				xAxis: [
					{
						type: 'category',
						boundaryGap: false,
						data: result.date,
					}
				],
				yAxis: [
					{
						type: 'value',
						axisLabel: {
							formatter: '{value} $'
						}
					}
				],
				series: [
					{
						symbol: "none",
						name: '花费',
						type: 'line',
						data: result.spend,
						markLine: {
							data: [
								{ type: 'average', name: '平均值' }
							]
						}
					},
					{
						symbol: "none",
						name: '收入',
						type: 'line',
						data: result.revenue,
						markLine: {
							data: [
								{ type: 'average', name: '平均值' }
							]
						}
					},
					{
						symbol: "none",
						name: '盈利',
						type: 'line',
						data: result.profitAndLoss,
						markLine: {
							data: [
								{ type: 'average', name: '平均值' }
							]
						}
					}
				]
			}
			require.config({
				paths: {
					echarts: 'http://echarts.baidu.com/build/dist'
				}
			});
			// 使用
			if (firstNum === false) {
				require(
					[
						'echarts',
						'echarts/chart/line' // 使用柱状图就加载bar模块，按需加载
					],
					function (ec) {
						// 基于准备好的dom，初始化echarts图表
						var myChart = ec.init(document.getElementById('show_chart'), "macarons");
						var option = optionInit;
						// 为echarts对象加载数据
						myChart.setOption(option);
					}
				);
				firstNum = true
			} else {
				require('echarts').init(document.getElementById('show_chart'), "macarons").setOption(optionInit);
			}
		}
	});

};

function ShowChart2() {
	result = char_data
	date = result.Date
	spend = result.Spend
	revenue = result.Revene
	var optionInit = {
		color: ['#FF0000', '#87cefa', '#33CC52', '#32cd32', '#6495ed'],
		backgroundColor: '#ffffff',
		title: {
			text: 'txtnation回本周期',
			subtext: ''
		},
		tooltip: {
			trigger: 'axis'
		},
		legend: {
			data: ['累计花费', '累计收入']
		},
		toolbox: {
			show: true,
			feature: {
				dataView: { show: true, readOnly: true },
				// magicType: { show: true, type: ['line', 'bar', 'stack', 'tiled'] },
				restore: { show: true },
				saveAsImage: { show: true }
			}
		},
		calculable: true,
		xAxis: [
			{
				type: 'category',
				boundaryGap: false,
				data: date,
			}
		],
		yAxis: [
			{
				type: 'value',
				axisLabel: {
					formatter: '{value} $'
				}
			}
		],
		series: [
			{
				symbol: "none",
				name: '累计花费',
				type: 'line',
				data: spend,
			},
			{
				symbol: "none",
				name: '累计收入',
				type: 'line',
				data: revenue,
			},
		]
	}
	require.config({
		paths: {
			echarts: 'http://echarts.baidu.com/build/dist'
		}
	});
	// 使用
	if (firstNum2 === false) {
		require(
			[
				'echarts',
				'echarts/chart/line' // 使用柱状图就加载bar模块，按需加载
			],
			function (ec) {
				// 基于准备好的dom，初始化echarts图表
				var myChart = ec.init(document.getElementById('show_chart_revenue'), "macarons");
				var option = optionInit;
				// 为echarts对象加载数据
				myChart.setOption(option);
			}
		);
		firstNum2 = true
	} else {
		require('echarts').init(document.getElementById('show_chart_revenue'), "macarons").setOption(optionInit);
	}

};

function ShowChart3() {
	result = char_data
	date = result.Date
	active = result.Active
	success = result.SuccessMt
	unsub = result.UnsubNum
	var optionInit = {
		color: ['#FF0000', '#87cefa', '#33CC52', '#32cd32', '#6495ed'],
		backgroundColor: '#ffffff',
		title: {
			text: 'txtnation留存变化趋势',
			subtext: ''
		},
		tooltip: {
			trigger: 'axis'
		},
		legend: {
			data: ['留存', '扣费次数', '退订']
		},
		toolbox: {
			show: true,
			feature: {
				dataView: { show: true, readOnly: true },
				// magicType: { show: true, type: ['line', 'bar', 'stack', 'tiled'] },
				restore: { show: true },
				saveAsImage: { show: true }
			}
		},
		calculable: true,
		xAxis: [
			{
				type: 'category',
				boundaryGap: false,
				data: date,
			}
		],
		yAxis: [
			{
				type: 'value',
				axisLabel: {
					formatter: '{value}'
				}
			}
		],
		series: [
			{
				symbol: "none",
				name: '留存',
				type: 'line',
				data: active,
			},
			{
				symbol: "none",
				name: '扣费次数',
				type: 'line',
				data: success,
			},
			{
				symbol: "none",
				name: '退订',
				type: 'line',
				data: unsub,
			},
		]
	}
	require.config({
		paths: {
			echarts: 'http://echarts.baidu.com/build/dist'
		}
	});
	// 使用
	if (firstNum3 === false) {
		require(
			[
				'echarts',
				'echarts/chart/line' // 使用柱状图就加载bar模块，按需加载
			],
			function (ec) {
				// 基于准备好的dom，初始化echarts图表
				var myChart = ec.init(document.getElementById('show_chart_subnum'), "macarons");
				var option = optionInit;
				// 为echarts对象加载数据
				myChart.setOption(option);
			}
		);
		firstNum3 = true
	} else {
		require('echarts').init(document.getElementById('show_chart_subnum'), "macarons").setOption(optionInit);
	}

};


function ShowChart4() {
	result = char_data1
	date = result.Date
	spend = result.Spend
	revenue = result.Revene
	var optionInit = {
		color: ['#FF0000', '#87cefa', '#33CC52', '#32cd32', '#6495ed'],
		backgroundColor: '#ffffff',
		title: {
			text: 'txtnation回本周期',
			subtext: ''
		},
		tooltip: {
			trigger: 'axis'
		},
		legend: {
			data: ['累计花费', '累计收入']
		},
		toolbox: {
			show: true,
			feature: {
				dataView: { show: true, readOnly: true },
				// magicType: { show: true, type: ['line', 'bar', 'stack', 'tiled'] },
				restore: { show: true },
				saveAsImage: { show: true }
			}
		},
		calculable: true,
		xAxis: [
			{
				type: 'category',
				boundaryGap: false,
				data: date,
			}
		],
		yAxis: [
			{
				type: 'value',
				axisLabel: {
					formatter: '{value} $'
				}
			}
		],
		series: [
			{
				symbol: "none",
				name: '累计花费',
				type: 'line',
				data: spend,
			},
			{
				symbol: "none",
				name: '累计收入',
				type: 'line',
				data: revenue,
			},
		]
	}
	require.config({
		paths: {
			echarts: 'http://echarts.baidu.com/build/dist'
		}
	});
	// 使用
	if (firstNum4 === false) {
		require(
			[
				'echarts',
				'echarts/chart/line' // 使用柱状图就加载bar模块，按需加载
			],
			function (ec) {
				// 基于准备好的dom，初始化echarts图表
				var myChart = ec.init(document.getElementById('show_chart_revenue1'), "macarons");
				var option = optionInit;
				// 为echarts对象加载数据
				myChart.setOption(option);
			}
		);
		firstNum2 = true
	} else {
		require('echarts').init(document.getElementById('show_chart_revenue1'), "macarons").setOption(optionInit);
	}

};

function ShowChart5() {
	result = char_data1
	date = result.Date
	active = result.Active
	success = result.SuccessMt
	unsub = result.UnsubNum
	var optionInit = {
		color: ['#FF0000', '#87cefa', '#33CC52', '#32cd32', '#6495ed'],
		backgroundColor: '#ffffff',
		title: {
			text: 'txtnation留存变化趋势',
			subtext: ''
		},
		tooltip: {
			trigger: 'axis'
		},
		legend: {
			data: ['留存', '扣费次数', '退订']
		},
		toolbox: {
			show: true,
			feature: {
				dataView: { show: true, readOnly: true },
				// magicType: { show: true, type: ['line', 'bar', 'stack', 'tiled'] },
				restore: { show: true },
				saveAsImage: { show: true }
			}
		},
		calculable: true,
		xAxis: [
			{
				type: 'category',
				boundaryGap: false,
				data: date,
			}
		],
		yAxis: [
			{
				type: 'value',
				axisLabel: {
					formatter: '{value}'
				}
			}
		],
		series: [
			{
				symbol: "none",
				name: '留存',
				type: 'line',
				data: active,
			},
			{
				symbol: "none",
				name: '扣费次数',
				type: 'line',
				data: success,
			},
			{
				symbol: "none",
				name: '退订',
				type: 'line',
				data: unsub,
			},
		]
	}
	require.config({
		paths: {
			echarts: 'http://echarts.baidu.com/build/dist'
		}
	});
	// 使用
	if (firstNum5 === false) {
		require(
			[
				'echarts',
				'echarts/chart/line' // 使用柱状图就加载bar模块，按需加载
			],
			function (ec) {
				// 基于准备好的dom，初始化echarts图表
				var myChart = ec.init(document.getElementById('show_chart_subnum1'), "macarons");
				var option = optionInit;
				// 为echarts对象加载数据
				myChart.setOption(option);
			}
		);
		firstNum3 = true
	} else {
		require('echarts').init(document.getElementById('show_chart_subnum1'), "macarons").setOption(optionInit);
	}

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
	$("#EveryDaySubQuality").hide();
	$("#UserComplaints").hide()
	$("#ChartEveryDay").hide()
});

$("#searchThird").click(function () {
	var start_time = document.getElementById("start_time_aff").value ? document.getElementById("start_time_aff").value : GetSevenDayDate();
	var end_time = document.getElementById("end_time_aff").value ? document.getElementById("end_time_aff").value : NowDate();

	data_aff = {
		"start_time": start_time,
		"end_time": end_time,
		"operator": $("#Opearator").val(),
		"aff_name": $("#Aff_Name option:selected").text(),
		"serverType": $("#service_type").val(),
		"clickType": $("#click_type option:selected").text()
	};
	getAffiliateData();
});



$("#searchMoAffDate").click(function () {
	var start_sub_date = document.getElementById("start_sub_date").value ? document.getElementById("start_sub_date").value : GetSevenDayDate();
	start_sub_date = start_sub_date.substr(0, 10);
	var end_sub_date = document.getElementById("end_sub_date").value ? document.getElementById("end_sub_date").value : NowDate();
	end_sub_date = end_sub_date.substr(0, 10);
	var start_date = document.getElementById("start_date").value ? document.getElementById("start_date").value : GetSevenDayDate();
	start_date = start_date.substr(0, 10);
	var end_date = document.getElementById("end_date").value ? document.getElementById("end_date").value : NowDate();
	end_date = end_date.substr(0, 10);


	subMoData = {
		"start_sub": start_sub_date,
		"end_sub": end_sub_date,
		"start_date": start_date,
		"end_date": end_date,
		"operator": $("#Quality_Opearator1").val(),
		"aff_name": $("#Quality_Aff_Name1 option:selected").text(),
		"service_type": $("#Quality_service_type1").val(),
		"clickType": $("#Quality_click_type1 option:selected").text(),
		"pub_id": $("#pud_select1 option:selected").text(),
	};
	// getMoDateData();
	getEerytimeQualitypage()
});


$("#searchEveryDaySubData").click(function () {
	var monthDate = document.getElementById("sale_month").value ? document.getElementById("sale_month").value : NowMoth();
	subeveryDayData = {
		"date": monthDate
	};
	GetEveryDaySubPage();
});

$("#searchComplaints").click(function () {
	var msisdn = document.getElementById("msisdn").value;
	ComplaintsData(msisdn)
});

$("#searchChart").click(function () {
	ShowChart()
});

//change nav page
$("#view_title").click(function () {
	$("#Quality").hide();
	$("#Subscriber").show();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#EveryDaySubQuality").hide();
	$("#UserComplaints").hide();
	$("#ChartEveryDay").hide()

	$("#back3").removeClass("background");
	$("#back4").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").removeClass("background");
	$("#back1").addClass("background");
	$("#back6").removeClass("background");
	$("#back7").removeClass("background");

});
$("#qm_title").click(function () {
	$("#Subscriber").hide();
	$("#Quality").show();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#EveryDaySubQuality").hide();
	$("#UserComplaints").hide()
	$("#ChartEveryDay").hide()

	$("#back1").removeClass("background");
	$("#back4").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").removeClass("background");
	$("#back3").addClass("background");
	$("#back6").removeClass("background");
	$("#back7").removeClass("background");

	QualityDatas = {
		"aff_name": "All",
		"operator": "All",
		"date": "2017-09-04"
	};
	// getQualitypage();
});
$("#aff_title").click(function () {
	$("#Subscriber").hide();
	$("#Quality").hide();
	$("#AffSubQuality").hide();
	$("#Everydaysubquality").hide();
	$("#UserComplaints").hide()
	$("#ChartEveryDay").hide()
	$("#Affiliate").show();

	$("#back3").removeClass("background");
	$("#back4").removeClass("background");
	$("#back1").removeClass("background");
	$("#back5").removeClass("background");
	$("#back6").removeClass("background");
	$("#back7").removeClass("background");
	$("#back2").addClass("background");


	data_aff = {
		"start_time": GetSevenDayDate(),
		"end_time": NowDate(),
		"operator": "All",
		"aff_name": "All",
		"serverType": "All"
	};
	// getAffiliateData();
});

$("#sub_qm_title").click(function () {
	$("#Subscriber").hide();
	$("#Quality").hide();
	$("#Affiliate").hide();
	$("#EveryDaySubQuality").hide();
	$("#AffSubQuality").show();
	$("#UserComplaints").hide()
	$("#ChartEveryDay").hide()


	$("#back3").removeClass("background");
	$("#back1").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").removeClass("background");
	$("#back6").removeClass("background");
	$("#back4").addClass("background");
	$("#back7").removeClass("background");


	var start_sub_date = document.getElementById("start_sub_date").value ? document.getElementById("start_sub_date").value : GetSevenDayDate();
	var end_sub_date = document.getElementById("end_sub_date").value ? document.getElementById("end_sub_date").value : NowDate();
	var start_date = document.getElementById("start_date").value ? document.getElementById("start_date").value : GetSevenDayDate();
	var end_date = document.getElementById("end_date").value ? document.getElementById("end_date").value : NowDate();
	subMoData = {
		"start_sub": start_sub_date,
		"end_sub": end_sub_date,
		"start_date": start_date,
		"end_date": end_date,
		"operator": $("#Opearator2").val(),
		"aff_name": $("#Aff_Name2 option:selected").text(),
		"service_type": $("#service_type2").val()
	};
	// getMoDateData();
});



$("#every_sub_qm").click(function () {
	$("#Subscriber").hide();
	$("#Quality").hide();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#UserComplaints").hide()
	$("#EveryDaySubQuality").show();
	$("#ChartEveryDay").hide()


	$("#back3").removeClass("background");
	$("#back4").removeClass("background");
	$("#back1").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").addClass("background");
	$("#back6").removeClass("background");
	$("#back7").removeClass("background");


	// GetEveryDaySubPage()
});

$("#Complaints").click(function () {
	$("#Subscriber").hide();
	$("#Quality").hide();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#UserComplaints").show();
	$("#EveryDaySubQuality").hide();
	$("#ChartEveryDay").hide()


	$("#back3").removeClass("background");
	$("#back4").removeClass("background");
	$("#back1").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").removeClass("background");
	$("#back6").addClass("background");
	$("#back7").removeClass("background");

	// ComplaintsData()
});

$("#Chart").click(function () {
	$("#Subscriber").hide();
	$("#Quality").hide();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#UserComplaints").hide();
	$("#EveryDaySubQuality").hide();
	$("#ChartEveryDay").show()


	$("#back3").removeClass("background");
	$("#back4").removeClass("background");
	$("#back1").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").removeClass("background");
	$("#back6").removeClass("background");
	$("#back7").addClass("background");

	// ShowChart()
});


// view page search
$("#query12").click(function () {
	viewDatas = {
		"operator": $("#telco_view").val(),
		"start_time": document.getElementById("start_time_view").value ? document.getElementById("start_time_view").value : GetSevenDayDate(),
		"end_time": document.getElementById("end_time_view").value ? document.getElementById("end_time_view").value : NowDate(),
		"aff_name": $("#affName").val(),
		"service_type": $("#serviceType").val(),
		"pubid": $("#pubId option:selected").text(),
		"clickType": $("#clickType option:selected").text()
	};

	$.ajax({
		url: 'http://pl.leadernethksp.com/aff_mt',
		type: 'GET',
		data: viewDatas,
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
		"operator": $("#Quality_Opearator").val(),
		"sub_date": time,
		"end_date": end_date,
		"serverType": $("#Quality_service_type").val(),
		"pub_id": $("#pud_select option:selected").text(),
		"clickType": $("#Quality_click_type option:selected").text()
	};
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



function changeMe1() {
	$.ajax({
		url: 'http://pl.leadernethksp.com/get_pubid?aff_name=' + $("#affName").val(),
		type: 'GET',
		dataType: "json",
		success: function (result) {
			var data = result.data;
			var aff_html = [];
			aff_html.push('<select class="selectpicker" id="pud_select_sub" data-live-search="true" title="Please select a lunch ..."><option>All</option>')
			$.each(data, function (o, p) {
				aff_html.push('<option>' + p + '</option>');
			});
			aff_html.push('</select>')
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
		url: 'http://pl.leadernethksp.com/get_pubid?aff_name=' + a,
		type: 'GET',
		dataType: "json",
		success: function (result) {
			var data = result.data;
			var aff_html = [];
			aff_html.push('<select class="selectpicker" id="pud_select" data-live-search="true" title="Please select a lunch ..."><option>All</option>')
			$.each(data, function (o, p) {
				aff_html.push('<option>' + p + '</option>');
			});
			aff_html.push('</select>')
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

function changeMeAffEverytime() {
	var a = $("#Quality_Aff_Name1").val();
	$.ajax({
		url: 'http://pl.leadernethksp.com/get_pubid?aff_name=' + a,
		type: 'GET',
		dataType: "json",
		success: function (result) {
			var data = result.data;
			var aff_html = [];
			aff_html.push('<select class="selectpicker" id="pud_select1" data-live-search="true" title="Please select a lunch ..."><option>All</option>')
			$.each(data, function (o, p) {
				aff_html.push('<option>' + p + '</option>');
			});
			aff_html.push('</select>')
			$("#ziqudao1").html(aff_html.join(""));
			$('.selectpicker').selectpicker({
				'selectedText': 'cat'
			});
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	});
}


function changeMe2() {
	var a = $("#Com_aff").val()
	$.ajax({
		url: 'http://pl.leadernethksp.com/get_pubid?aff_name=' + a,
		type: 'GET',
		dataType: "json",
		success: function (result) {
			var data = result.data;
			var aff_html = [];
			aff_html.push('<select class="selectpicker" id="Com_pub_select" data-live-search="true" title="Please select a lunch ..."><option>All</option>')
			$.each(data, function (o, p) {
				aff_html.push('<option>' + p + '</option>');
			});
			aff_html.push('</select>')
			$("#Com_ziqudao").html(aff_html.join(""));
			$('.selectpicker').selectpicker({
				'selectedText': 'cat'
			});
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	});
}

function changeMe3() {
	$.ajax({
		url: 'http://pl.leadernethksp.com/get_pubid?aff_name=' + $("#Aff_Name_char").val(),
		type: 'GET',
		dataType: "json",
		success: function (result) {
			var data = result.data;
			var aff_html = [];
			aff_html.push('<select class="selectpicker" id="Pud_select_char" data-live-search="true" title="Please select a lunch ..."><option>All</option>')
			$.each(data, function (o, p) {
				aff_html.push('<option>' + p + '</option>');
			});
			aff_html.push('</select>')
			$("#ziqudao_char").html(aff_html.join(""));
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
	language: 'zh-CN',
	format: 'yyyy-mm-dd',
	todayBtn: 1,
	autoclose: 1,
});



function downloadTable1Excal(aff_name, PubId, service_type, operator, start_sub, end_sub) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>AffName:" + aff_name + "  |PubId:" + PubId + "   |ServiceType:" + service_type + "  |Operator:" + operator + "   |Start Sub Date: " + start_sub + " |End Sub Date: " + end_sub + "</h4>" + document.getElementById("table1").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文件内容的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down1");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "扣费数量.xls";
}

function downloadTable2Excal(aff_name, service_type, operator, start_sub, end_sub) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>AffName:" + aff_name + "   |ServiceType:" + service_type + "  |Operator:" + operator + "   |Start Sub Date: " + start_sub + " |End Sub Date: " + end_sub + "</h4>" + document.getElementById("table2").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文件内容的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down2");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "渠道质量.xls";
}


function downloadTable3Excal(aff_name, PubId, service_type, operator, start_sub, end_sub) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>AffName:" + aff_name + "   |PubId:" + PubId + "      |ServiceType:" + service_type + "  |Operator:" + operator + "   |Start Sub Date: " + start_sub + " |End Sub Date: " + end_sub + "</h4>" + document.getElementById("table3").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文件内容的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down3");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "持续扣费情况.xls";
}


function downloadTable4Excal(aff_name, service_type, operator, start_sub, end_sub, start, end) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>AffName:" + aff_name + "     |ServiceType:" + service_type + "  |Operator:" + operator + "   |Start Sub Date: " + start_sub + " |End Sub Date: " + end_sub + "   |Search Start Date: " + start + "   |Search End Date: " + end + "</h4>" + document.getElementById("table4").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down4");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "订阅留存.xls"
}

function downloadTable5Excal(Month) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>" + Month + "</h4>" + document.getElementById("table5").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down5");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "订阅留存.xls"
}


$("#sale_month").datepicker({
	changeMonth: true,
	changeYear: true,
	dateFormat: 'mm-yy',
	showButtonPanel: true,
	monthNamesShort: ['01', '02', '03', '04', '05', '06', '07', '08', '09', '10', '11', '12'],
	closeText: '选择',
	currentText: '本月',
	isSelMon: 'true',
	onClose: function (dateText, inst) {
		var month = +$("#ui-datepicker-div .ui-datepicker-month :selected").val() + 1,
			year = $("#ui-datepicker-div .ui-datepicker-year :selected").val();
		if (month < 10) {
			month = '0' + month;
		}
		this.value = year + '-' + month;
		if (typeof this.blur === 'function') {
			this.blur();
		}
	}
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
	months = (myDate.getMonth() + 1).toString()
	days = myDate.getDate()
	if (months.length === 1){
		months = "0" + months
	}
	if (days.length === 1){
		days = "0" + days
	}
	dateTemp = myDate.getFullYear() + "-" + months + "-" + days;
	return dateTemp
}

function NowMoth() {
	var myDate = new Date(); //获取今天日期
	var dateTemp;
	months = (myDate.getMonth() + 1).toString()
	days = myDate.getDate()
	if (months.length === 1){
		months = "0" + months
	}
	dateTemp = myDate.getFullYear() + "-" + months;
	return dateTemp
}

//投诉页面
$("#com_submit").click(function () {
	var com_data = {
		"Msisdn": $("#com_msisdn").val(),
		"Email": $("#com_email").val(),
		"UserName": $("#com_user").val(),
		"EquipmentModel": $("#com_equi").val(),
		"GuiltyAffName": $("#com_aff").val(),
		"GuiltyPubid": $("#com_pub").val(),
		"Description": $("#com_des").val(),
		"Level": $("#Com_level").val(),
		"DealWithTime": $("#Com_time").val()
	};
	$.ajax({
		url: 'http://pl.leadernethksp.com/addComplaint',
		type: 'POST',
		data: JSON.stringify(com_data),
		contentType: "application/json",
		dataType: "json",
		success: function (result) {
			alert(result.message)
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	})
})

$("#com_search").click(function () {
	$("#com_info").addClass("noexit")
	var com_data = {
		"start": $("#start_com_date").val().substr(0, 10),
		"end": $("#end_com_date").val().substr(0, 10),
		"service_type": $("#Com_service").val(),
		"operator": $("#Com_operator").val(),
		"aff_name": $("#Com_aff").val(),
		"pub_id": $("#Com_pub_select").val(),
		"clickType": $("#Com_click").val(),
		"level": $("#Com_level_sel").val()
		// "msisdn":"asdsd",
	};
	$.ajax({
		url: 'http://pl.leadernethksp.com/get/complaint/data',
		type: 'GET',
		data: com_data,
		dataType: "json",
		success: function (result) {
			var data = result;
			var aff_html = [];
			$.each(data.data, function (i, c) {
				aff_html.push('<tr><td rowspan=' + c.ComplaintList.length + '>' + c.Date + '</td>');
				$.each(c.ComplaintList, function (i, c) {
					if (i == 0) {
						aff_html.push('<td>' + c.Msisdn + '</td><td>' + c.Operator + '</td><td>' + c.UserName + '</td><td>' + c.Email + '</td><td>' + c.AffName + '</td><td>' + c.PubId + '</td><td>' + c.ClickId + '</td><td>' + c.ClickType + '</td><td>' + c.PostbackStatus + '</td><td>' + c.MtNum + '</td><td>' + c.Amount + '</td><td>' + c.RackingCode + '</td><td>' + c.Subtime + '</td><td>' + c.Unsubtime + '</td><td>' + c.SubId + '</td><td>' + c.EquipmentModel + '</td><td>' + c.GuiltyAffName + '</td><td>' + c.GuiltyPubid + '</td><td>' + c.Description + '</td>');
					} else {
						aff_html.push('<tr><td>' + c.Msisdn + '</td><td>' + c.Operator + '</td><td>' + c.UserName + '</td><td>' + c.Email + '</td><td>' + c.AffName + '</td><td>' + c.PubId + '</td><td>' + c.ClickId + '</td><td>' + c.ClickType + '</td><td>' + c.PostbackStatus + '</td><td>' + c.MtNum + '</td><td>' + c.Amount + '</td><td>' + c.RackingCode + '</td><td>' + c.Subtime + '</td><td>' + c.Unsubtime + '</td><td>' + c.SubId + '</td><td>' + c.EquipmentModel + '</td><td>' + c.GuiltyAffName + '</td><td>' + c.GuiltyPubid + '</td><td>' + c.Description + '</td>');
					}
				});
				aff_html.push('</tr>')
			});

			var aff_total = [];

			$.each(data.totalData, function (i, c) {
				aff_total.push('<tr>');
				aff_total.push('<td>' + c.AffName + '</td><td>' + c.ComplaintNum + '</td><td>' + c.PostbackNum + '</td><td>' + c.MtNum + '</td><td>' + c.Amount + '</td><td>' + c.Level_1 + '</td><td>' + c.Level_2 + '</td><td>' + c.Level_3 + '</td>');
				aff_total.push('</tr>')
			});
			$("#com_table_sel_data").html(aff_html.join(""));
			$("#com_table_sel_data_first").html(aff_total.join(""));
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	})
})


$("#table_char").click(function () {
	$("#show_table").removeClass("noexit")
	$("#show_chart_revenue").addClass("noexit")
	$("#show_chart_subnum").addClass("noexit")
	$("#daochu3").removeClass("noexit")
})
$("#huiben_char").click(function () {
	$("#show_chart_revenue").removeClass("noexit")
	$("#show_chart_subnum").addClass("noexit")
	$("#show_table").addClass("noexit")
	$("#daochu3").addClass("noexit")
	ShowChart2()
})
$("#liucun_char").click(function () {
	$("#show_chart_subnum").removeClass("noexit")
	$("#show_table").addClass("noexit")
	$("#show_chart_revenue").addClass("noexit")
	$("#daochu3").addClass("noexit")
	ShowChart3()
})


$("#table_char1").click(function () {
	$("#show_table1").removeClass("noexit")
	$("#show_chart_revenue1").addClass("noexit")
	$("#show_chart_subnum1").addClass("noexit")
	$("#daochu4").removeClass("noexit")
})
$("#huiben_char1").click(function () {
	$("#show_chart_revenue1").removeClass("noexit")
	$("#show_chart_subnum1").addClass("noexit")
	$("#show_table1").addClass("noexit")
	$("#daochu4").addClass("noexit")
	ShowChart4()
})
$("#liucun_char1").click(function () {
	$("#show_chart_subnum1").removeClass("noexit")
	$("#show_table1").addClass("noexit")
	$("#show_chart_revenue1").addClass("noexit")
	$("#daochu4").addClass("noexit")
	ShowChart5()
})