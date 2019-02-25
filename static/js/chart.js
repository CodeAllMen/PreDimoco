var date, spend, profitAndLoss, revenue;
$.ajax({
    url: 'http://172.16.60.13:8087/day/chart',
    type: 'GET',
    data: { 'startDate': "2018", "endDate": "2019" },
    dataType: 'json',
    success: function (result) {
        date = result.date
        spend = result.spend
        profitAndLoss = result.profitAndLoss
        revenue = result.revenue
    }
});
require.config({
    paths: {
        echarts: 'http://echarts.baidu.com/build/dist'
    }
});
// 使用
require(
    [
        'echarts',
        'echarts/chart/line' // 使用柱状图就加载bar模块，按需加载
    ],
    function (ec) {
        // 基于准备好的dom，初始化echarts图表
        var myChart = ec.init(document.getElementById('show_chart'));
        var option = {
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
                    magicType: { show: true, type: ['line'] },
                    restore: { show: true },
                    saveAsImage: { show: true }
                }
            },
            calculable: true,
            xAxis: [
                {
                    type: 'category',
                    boundaryGap: false,
                    data: date
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
                    name: '花费',
                    type: 'line',
                    data: spend,
                    markLine: {
                        data: [
                            { type: 'average', name: '平均值' }
                        ]
                    }
                },
                {
                    name: '收入',
                    type: 'line',
                    data: revenue,
                    markLine: {
                        data: [
                            { type: 'average', name: '平均值' }
                        ]
                    }
                },
                {
                    name: '盈利',
                    type: 'line',
                    data: profitAndLoss,
                    markLine: {
                        data: [
                            { type: 'average', name: '平均值' }
                        ]
                    }
                }
            ]
        };
        // 为echarts对象加载数据
        myChart.setOption(option);
        myChart.setTheme("macarons")
    }
);