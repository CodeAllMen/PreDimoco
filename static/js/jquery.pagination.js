(function ($) {

    $.initPagination = function (info) {
        var $div = $(this);

        var pager = {
            $first: null,//首页按钮
            $prev: null,//上一页按钮
            $nav: null,//导航区按钮
            $next: null,//下一页按钮
            $last: null,//末页按钮
            $goNum: null,//翻页文本输入框
            $go: null,//翻页按钮
            $recordCount: null,//总行数
            $pageIndex: null,//当前页号
            $pageCount: null,//总页数
            $pageSize: null,//页码数
            Init: function () { //初始化控件 缓存
                $div.find("[name]").each(function () {
                    var name = $(this).attr("name");
                    if (name && name.length) {
                        pager["$" + name] = $(this);
                    }
                });
                pager.SetPageUI();
            },
            //创建导航
            CreateNavi: function () {
                //页号导航
                if (!pager.$nav) return;

                var n1 = "<a class=\"navi\">{text}</a>";
                var navi = pager.$nav;
                navi.empty();
                var i; var temp = ""; var s1; var s2;

                var naviCount2 = (info.naviCount % 2) == 1 ? info.naviCount / 2 | 0 : info.naviCount / 2;

                if ((info.thisPageIndex - naviCount2) < 1) {
                    s1 = 1; s2 = info.naviCount;
                } else if ((info.thisPageIndex + naviCount2) >= info.pageCount) {
                    s1 = info.pageCount - info.naviCount + 1;
                    s2 = info.pageCount;
                    if (s1 < 1) s1 = 1;
                } else {
                    s1 = info.thisPageIndex - naviCount2;
                    s2 = info.thisPageIndex + naviCount2;
                    if (s2 > info.pageCount) s2 = info.pageCount;
                }
                if (s2 > info.pageCount) s2 = info.pageCount;

                if (s1 >= 2) {
                    temp = n1.replace("{text}", 1);
                    navi.append(temp);
                }
                if (s1 >= 3) {
                    temp = n1.replace("{text}", "…");
                    navi.append(temp);
                }

                for (i = s1; i <= s2; i++) { //当前页
                    temp = n1.replace("{text}", i == info.thisPageIndex ? "<font style=\"color:#999\">" + i + "</font>" : i);
                    navi.append(temp);
                }

                if (s2 <= info.pageCount - 2) {
                    temp = n1.replace("{text}", "…");
                    navi.append(temp);
                }
                if (s2 <= info.pageCount - 1) {
                    temp = n1.replace("{text}", info.pageCount);
                    navi.append(temp);
                }

                //注册事件
                var naviA = navi.find("a");
                naviA.each(function (j) {
                    $(this).on('click', function () {
                        info.beforePageIndex = info.thisPageIndex;
                        if (this.innerHTML == "…") {
                            var a1 = naviA[j - 1];
                            if (a1.innerHTML == "1") {
                                //前面的，取后面的作为页号
                                a1 = naviA[j + 1];
                                info.thisPageIndex = $(a1).text() * 1 - 1;
                            } else {
                                //后面的，取前面的作为页号
                                a1 = naviA[j - 1];
                                info.thisPageIndex = $(a1).text() * 1 + 1;
                            }

                        } else {
                            info.thisPageIndex = $(this).text();
                        }
                        //页码未变更则不触发事件
                        if (info.beforePageIndex == info.thisPageIndex) return;
                        pager.PageChanged();
                    });
                });
            },
            //设置UI
            SetPageUI: function () {
                if (info.pageCount == -1) {
                    //计算总页数
                    info.pageCount = info.recordCount % info.pageSize === 0 ? info.recordCount / info.pageSize : parseInt(info.recordCount / info.pageSize) + 1;
                }

                pager.$recordCount && pager.$recordCount.html(info.recordCount);
                pager.$pageIndex && pager.$pageIndex.html(info.thisPageIndex);
                pager.$pageCount && pager.$pageCount.html(info.pageCount);
                pager.$pageSize && pager.$pageSize.html(info.pageSize);

                pager.$goNum && pager.$goNum.val(info.thisPageIndex);

                $div.find("span").css("cursor", "pointer");

                switch (parseInt(info.thisPageIndex)) {
                    case 1:
                        //第一页
                        pager.$first && pager.$first.hide().css("cursor", "default");
                        pager.$prev && pager.$prev.hide().css("cursor", "default");

                        pager.$next && pager.$next.show().css("cursor", "pointer");
                        pager.$last && pager.$last.show().css("cursor", "pointer");
                        break;
                    case info.pageCount:
                        //最后一页
                        pager.$first && pager.$first.show().css("cursor", "pointer");
                        pager.$prev && pager.$prev.show().css("cursor", "pointer");

                        pager.$next && pager.$next.hide().css("cursor", "default");
                        pager.$last && pager.$last.hide().css("cursor", "default");
                        break;
                    default:
                        //不是第一页、最后一页
                        pager.$first && pager.$first.show().css("cursor", "pointer");
                        pager.$prev && pager.$prev.show().css("cursor", "pointer");

                        pager.$next && pager.$next.show().css("cursor", "pointer");
                        pager.$last && pager.$last.show().css("cursor", "pointer");
                        break;
                }
            },
            //设置缺省设置
            SetDefaultInfo: function (pSet) {
                var defaultSetting = {
                    recordCount: -1,        //总记录数
                    pageSize: 20,           //一页记录数
                    pageCount: -1,          //总页数，不用设置，根据上两个条件自动计算
                    thisPageIndex: 1,       //当前页号
                    beforePageIndex: 1,     //上一次的页号
                    //lastPageIndex: 0,     //最后一页的页号，自动计算

                    naviCount: 7,            //页号导航的数量
                    isUseRecordCount: true,   //是否使用缓存的总记录数

                    //分页控件模板的路径和文件名
                    templetPath: "pager",
                };
                return $.extend(defaultSetting, pSet);
            },
            //注册翻页的事件，不包括页号导航
            RegPagerEvent: function () {
                pager.$first && pager.$first.on('click', function () {
                    if (info.thisPageIndex != 1) {
                        addEvent("a");
                    }
                });

                pager.$prev && pager.$prev.on('click', function () {
                    if (info.thisPageIndex != 1) {
                        addEvent("-");
                    }
                });

                pager.$next && pager.$next.on('click', function () {
                    if (info.thisPageIndex != info.pageCount) {
                        addEvent("+");
                    }
                });

                pager.$last && pager.$last.on('click', function () {
                    if (info.thisPageIndex != info.pageCount) {
                        addEvent("z");
                    }
                });

                pager.$go && pager.$go.on('click', function () {
                    addEvent("go");
                });

                var addEvent = function (kind) {
                    info.beforePageIndex = info.thisPageIndex;
                    switch (kind) {
                        case "a": info.thisPageIndex = 1; break; //首页 
                        case "z": info.thisPageIndex = info.pageCount; break; //末页 
                        case "+": info.thisPageIndex = info.thisPageIndex * 1 + 1; break; //下页 
                        case "-": info.thisPageIndex = info.thisPageIndex * 1 - 1; break; //上页 
                        case "go": info.thisPageIndex = pager.$goNum.val(); break; //go 
                    }
                    pager.PageChanged();
                };
            },
            //页面变更事件
            PageChanged: function (pIndex) {
                if (pIndex) info.thisPageIndex = pIndex;

                //如果页码错误 则中断
                if (isNaN(info.thisPageIndex)) {
                    alert("页码不正确");
                    return;
                }

                info.thisPageIndex = parseInt(info.thisPageIndex);

                //判断范围
                if (info.thisPageIndex > info.pageCount) info.thisPageIndex = info.pageCount;
                if (info.thisPageIndex < 1) info.thisPageIndex = 1;

                info.onPageChange && info.onPageChange.call(info, info.thisPageIndex);

                pager.SetPageUI(); //重新设置UI
                pager.CreateNavi();

                //克隆
                //var tmpdiv = info.pageTurnDivIDs.split(",");
                //var pageHtml = $("#" + tmpdiv[0]).clone(true);
                //for (var i = 1; i < tmpdiv.length; i++) {
                //    $("#" + tmpdiv[i]).html(pageHtml);
                //}
            }
        };

        //设置默认值
        info = pager.SetDefaultInfo(info);
        //如果模板没有自定义 则自动填充
        var html = $div.html();
        if (!/\w+/.test(html)) {
            html = '第<span name="pageIndex" style="color: Red;">1</span>/<span name="pageCount" style="color: Red;">11</span>页</span><span style="cursor: pointer;">    每页<span name="pageSize" style="color: Red;">10</span>条记录    </span><span name="first" class="disabled" style="cursor: default;">首页</span><span name="prev" class="disabled" style="cursor: default;">上一页  </span><span name="nav" id="nav" style="cursor: pointer;"><a class="navi"><span style="color: #999">1</span></a></span><span name="next" id="next" style="cursor: pointer;">下一页  </span><span name="last" class="" style="cursor: pointer;">末页    </span><input name="goNum" class="cssTxt" type="text" size="1"><span name="go" style="cursor: pointer;">   GO</span>';
            $div.html(html);
        }
        //增加样式
        if (!$div.hasClass("pager")) $div.addClass("pager");
        //初始化 缓存控件
        pager.Init();
        //注册事件（仅限上一页、下一页、首页、末页，不包括页号导航）
        pager.RegPagerEvent();
        //创建导航 并增加导航事件
        pager.CreateNavi();
        this.pager = pager;

        return this;
    }

    $.fn.pagination = function (pagerInfo) {
        this.each(function () {
            $.initPagination.call(this, pagerInfo);
        });
    }

    $.fn.PageChanged = function (pageIndex) {
        this[0].pager && this[0].pager.PageChanged.call(this, pageIndex);
    };
}(jQuery));