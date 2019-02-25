/**
 * Created by JingRu on 2017/8/31.
 */
$(function(){
    //首页图片的渲染
      index();
    gameLogin();
    //阻止事件冒泡
    function stopPropagation(e) {
        if (e.stopPropagation)
            e.stopPropagation();
        else
            e.cancelBubble = true;
    }

    //点击页面其他地方，导航隐藏
    $(document).bind('click',function(){
        //$(".nav-dropdown").css("display","none");
        $(".nav-dropdown").hide(500);
    });
    /*点击导航，导航不隐藏
    $(".nav-dropdown").click(function (e) {
        stopPropagation(e);
    });*/
    //点击games标签，导航显示
    $('#games').bind('click',function(e){
       // $(".nav-dropdown").css("display","block");
        $(".nav-dropdown").toggle(500);
        stopPropagation(e);
        $.ajax({
            url: 'http://127.0.0.1:8083/gettag',
            type: 'GET',
            // data: chargeViewDatas,
            dataType: "json",
            success: function(result){
                var data = result;
                // console.log(data);
                var chargeviewtb1 =[];
                $.each(data.data, function(i, c) {
                    var a= c.Id;
                    chargeviewtb1.push("<div class='col-sm-4 col-md-4 noPadding'><li><a href='#'  class='a-font' onclick='gameMore("+ c.Id+")' >"+ c.Tag_Name+"</a></li></div>");
                });
                $(".appendLi").html(chargeviewtb1.join(""));
            }.bind(this),
            error: function(error){
                console.log(error);
            }
        });
    });
/*登录按钮点击,第一版需要注册的
    $("#userLogin").click(function(){
        var name=$("#userNameOne").val();
        var password=$("#userPasswordOne").val();
        var userData={"name":name,"password":password};
        $.ajax({
            url: 'http://172.16.20.27:8085/login',
            type: 'post',
            data: JSON.stringify(userData),
            dataType: "json",
            success: function(result){
                var data=result;
                if(data.code=="0"){
                    alert(data.message);
                }else if(data.code=="1"){
                    var dataTwo=data.data;
                    var b='<div class="signTwo"><a href="#" class="btn showAll smallScreen" onclick="personal()">'+
                        ' <i class="icon icon-user"></i><span id="user_name">'+dataTwo.User_Name+'</span><input type="hidden" id="user_email" value="'+dataTwo.Email+'" /></a>'+
                        '<div class="smallPTwo modal-bodyText line"><i class="icon icon-money"></i>'+
                        '<div class="smallPTwo paddingTop"><span class="moneyFont">×</span>'+
                        '<lable class="moneyFont">'+dataTwo.Money+'</lable></div></div></div>'
                    $("#sign").html(b);
                    $(".close").click();

                }
            },
            error: function(error){
                console.log(error);
            }
        });
    });*/
//登录按钮新版测试，不需要注册
    $("#checkFirstLogin").click(function(){
        alert("124")
        var name=$("#userNameOne").val();
        var password=$("#userPasswordOne").val();
        if(name==""||name==null){
            alert("UserName cannot be empty!");
        }else if(password==""||password==null){
            alert("password name cannot be empty!");
        }else {
            var userData = {"name": name, "password": password};
            $.ajax({
                url: 'http://127.0.0.1:8083/login',
                type: 'post',
                data: JSON.stringify(userData),
                dataType: "json",
                async: false,
                success: function (result) {
                    var data = result;
                    if (data.code == "0") {
                        alert("UserName or password incorrect!");
                    } else if (data.code == "1") {
                        var dataTwo = data.data;
                        var b = '<div class="signTwo"><a href="#" class="btn showAll smallScreen" onclick="personal()">' +
                            ' <i class="icon icon-user"></i><span id="user_name">' + dataTwo.User_Name + '</span><input type="hidden" id="user_email" value="' + dataTwo.Email + '" /></a>' +
                            '<div class="smallPTwo modal-bodyText line"><i class="icon icon-money"></i>' +
                            '<div class="smallPTwo paddingTop"><span class="moneyFont">×</span>' +
                            '<lable class="moneyFont">' + dataTwo.Money + '</lable></div></div></div>'
                        $("#sign").html(b);
                        $(".close").click();

                    } else if (data.code == "2") {
                        $("#checkFirstLogin").attr("href", "#modal-container-801870");
                        $("#userNameTwo").val(name);
                    }
                },
                error: function (error) {
                    console.log(error);
                }
            });
        };
    })
/* 新的类似于注册，设置新密码和邮箱的弹出框*/
    $("#setInformation").click(function(){
        var userNameTwo=$("#userNameTwo").val();
        var passwordOne=$("#passwordOne").val();
        var passwordTwo=$("#passwordTwo").val();
        var email=$("#emailOne").val();
        var reg = /^([0-9A-Za-z\-_\.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g;
        if(passwordOne==""||passwordTwo==""){
            alert("Password cannot be empty!")
        }else if(passwordOne!=passwordTwo){
            alert("The two mailbox must be consistent");
        }else if(!reg.test(email)){
            alert("Please fill in the correct mailbox!");
        }else{
            var userData={"passwd":passwordOne,"email":email,"userName":userNameTwo};
            $.ajax({
                url: 'http://127.0.0.1:8083/register',
                type: 'post',
                data: JSON.stringify(userData),
                dataType: "json",
                success: function(result){
                    var data=result;
                    // console.log(data);
                    if(data.code=="0"){
                        alert(data.message);
                    }else if(data.code=="1"){
                        var dataTwo=data.data;
                        var b=',<div class="signTwo"><a href="#" class="btn showAll smallScreen" onclick="personal()">'+
                            ' <i class="icon icon-user"></i>'+dataTwo.User_Name+'</a>'+
                            '<div class="smallPTwo modal-bodyText line"><i class="icon icon-money"></i>'+
                            '<div class="smallPTwo paddingTop"><span class="moneyFont">×</span>'+
                            '<lable class="moneyFont">'+dataTwo.Money+'</lable></div></div></div>'
                        $("#sign").html(b);
                        $(".close").click();
                    }
                },
                error: function(error){
                    console.log(error);
                }
            });
        }
    });
/*注册按钮点击
    $("#signUp").click(function(){
        var name=$("#userNameTwo").val();
        var password=$("#userPasswordTwo").val();
        var email=$("#userEmail").val();
        var emailTwo=$("#userEmailTwo").val();
        var reg = /^([0-9A-Za-z\-_\.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g;
        if(name==""){
            alert("User name cannot be empty!")
        }else if(!reg.test(email)){
            alert("Please fill in the correct mailbox!");
        }else if(email!=emailTwo){
            alert("The two mailbox must be consistent");
        }else{
            var userData={"UserName":name,"Email":email,"Passwd":password};
            $.ajax({
                url: 'http://172.16.20.27:8085/register',
                type: 'post',
                data: JSON.stringify(userData),
                dataType: "json",
                success: function(result){
                    var data=result;
                   // console.log(data);
                    if(data.code=="0"){
                        alert(data.message);
                    }else if(data.code=="1"){
                        var dataTwo=data.data;
                        var b=',<div class="signTwo"><a href="#" class="btn showAll smallScreen" onclick="personal()">'+
                              ' <i class="icon icon-user"></i>'+dataTwo.User_Name+'</a>'+
                              '<div class="smallPTwo modal-bodyText line"><i class="icon icon-money"></i>'+
                              '<div class="smallPTwo paddingTop"><span class="moneyFont">×</span>'+
                              '<lable class="moneyFont">'+dataTwo.Money+'</lable></div></div></div>'
                        $("#sign").html(b);
                        $(".close").click();
                    }
                },
                error: function(error){
                    console.log(error);
                }
            });
        }
    })*/
//home列表的点击
    $("#home").click(function(){
        index();
    })

    //试玩游戏
    $("#demoGame").click(function(){
       var a=$("#game_idDemo").val();
       var userData={"game_id":a,"type":"demo"};
        $.ajax({
            url: 'http://127.0.0.1:8083/buygame',
            type: 'post',
            data: JSON.stringify(userData),
            dataType: "json",
            success: function(result){
                var data=result;
                if(data.code=="0"){
                    alert("每天有3次试玩机会，您的试玩机会已用完");
                }else if(data.code=="1"){
                    window.location.href=data.data;
                }
            },
            error: function(error){
                console.log(error);
            }
        });
    });
    //购买游戏
    $("#purchase").click(function(){
        var a=$("#game_idDemo").val();
        var userData={"game_id":a,"type":"buy"};
        $.ajax({
            url: 'http://127.0.0.1:8083/buygame',
            type: 'post',
            data: JSON.stringify(userData),
            dataType: "json",
            success: function(result){
                var data=result;
                window.location.href=result.data
            },
            error: function(error){
                console.log(error);
            }
        });
    });
})
//进入更多游戏页面
function gameMore(id){
    $.ajax({
        url: 'http://127.0.0.1:8083/taggame?tag_id='+id,
        type: 'GET',
        // data: chargeViewDatas,
        dataType: "json",
        success: function(result){
            var data = result;
           // console.log(data);
            var chargeviewtb1 =[];
            var dataTwo = data.data.Game_list;
             if(dataTwo!=null) {
                 var a = "<div class='container bodyDiv'><header>" +
                     "<span class='h4' >" + data.data.Tag_name + "</span><div class='moreButton'><a href='#' onclick='index()' class='btn showAll' > BACK </a></div>" +
                     "</header><div class='icon icon-line'></div><section><div class='col-sm-12'><div class='row row-top'><div class='col-sm-4 noPadding padding-top contentOne' >" +
                     "<img class='imgOne' src='" + dataTwo[0].Image_url + "'/><div class='txt'><h6>" + dataTwo[0].Game_name + "</h6><a href='#' class='btn playNow' onclick='gameDemo(" + dataTwo[0].Game_id + ")' >Play Now</a>" +
                     "</div></div>";
                 var b = "";
                 for (var j = 1; j < dataTwo.length; j = j + 4) {
                     var jOne = "", jTwo = "", jThree = "", jFour = "";
                     if (b == "") {
                         b = "<div class='col-sm-4 noPadding'>";
                         if (dataTwo[j] != undefined) {
                             jOne = "<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j].Image_url + "'  /><div class='txtTwo'><h6>"+ dataTwo[j].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j].Game_id+")'>Play Now</a>"+
                                 "</div></div>";
                         }
                         if (dataTwo[j + 1] != undefined) {
                             jTwo = "<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j + 1].Image_url + "'  /><div class='txtTwo'><h6>"+ dataTwo[j+1].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j+1].Game_id+")'>Play Now</a>"+
                                 "</div></div>";
                         }
                         if (dataTwo[j + 2] != undefined) {
                             jThree = "<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j + 2].Image_url + "'  /><div class='txtTwo'><h6>"+ dataTwo[j+2].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j+2].Game_id+")'>Play Now</a>"+
                                 "</div></div>";
                         }
                         if (dataTwo[j + 3] != undefined) {
                             jFour = "<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j + 3].Image_url + "'  /><div class='txtTwo'><h6>"+ dataTwo[j+3].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j+3].Game_id+")'>Play Now</a>"+
                                 "</div></div>";
                         }
                         b = b + jOne + jTwo + jThree + jFour + "</div>";
                     } else {
                         if (dataTwo[j] != undefined) {
                             jOne = "<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img src='" + dataTwo[j].Image_url + "'  class='imgTwo' /><div class='txtTwo'><h6>"+ dataTwo[j].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j].Game_id+")'>Play Now</a>"+
                                 "</div></div>";
                         }
                         if (dataTwo[j + 1] != undefined) {
                             jTwo = "<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j + 1].Image_url + "' /><div class='txtTwo'><h6>"+ dataTwo[j+1].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j+1].Game_id+")'>Play Now</a>"+
                                 "</div></div>";
                         }
                         if (dataTwo[j + 2] != undefined) {
                             jThree = "<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j + 2].Image_url + "'/><div class='txtTwo'><h6>"+ dataTwo[j+2].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j+2].Game_id+")'>Play Now</a>"+
                                 "</div></div>";
                         }
                         if (dataTwo[j + 3] != undefined) {
                             jFour = "<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j + 3].Image_url + "' /><div class='txtTwo'><h6>"+ dataTwo[j+3].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j+3].Game_id+")'>Play Now</a>"+
                                 "</div></div>";
                         }
                         b = b + "<div class='col-sm-4 noPadding'>" + jOne + jTwo + jThree + jFour + "</div>";

                     }
                 }
                 var d = "</div></div></section></div>";
                 chargeviewtb1.push(a + b + d);
                 $("#body").html(chargeviewtb1.join(""));
                 hover();
             }
        }.bind(this),
        error: function(error){
            console.log(error);
        }
    });
}
function gameDemo(id){
    $.ajax({
        url: 'http://127.0.0.1:8083/play?game_id='+id,
        type: 'GET',
        // data: chargeViewDatas,
        dataType: "json",
        success: function(result){
            var data = result;
            var chargeviewtb1 =[];
            var dataTwo = data.data.Small_img;
            if(dataTwo!=null) {
                //console.log(dataTwo);
                var b ='<div class="container bodyDiv"><header><span class="h4">Game Details</span><div class="moreButton"><a href="#" onclick="index()" class="btn showAll">< Back </a></div>'+
                        '</header><div class="icon icon-line"></div><section><div class="col-sm-12 col-xs-12"><div class="row row-top"><div class="col-sm-12 col-xs-12 noPadding padding-top">'+
                        '<img src="'+data.data.Big_img+'" class="img-Three" /></div></div><div class="row row-top"><div class="col-sm-4 noPadding"> <img src="'+dataTwo[0]+'"  class="img" />'+
                        '</div><div class="col-sm-4 noPadding "><img src="'+dataTwo[1]+'"  class="img" /></div><div class="col-sm-4 noPadding"><img src="'+dataTwo[2]+'"  class="img" />'+
                        '</div></div></div><div class="col-sm-12"><div class="col-sm-6 col-xs-12"><h4>'+data.data.Game_name+'</h4></div> <div class="col-sm-6 col-xs-10 priceText moneyPadding paddingBottom">'+
                        '<div class="smallPTwo moneyFont moneySize moneyPadding">×'+data.data.Price+'</div><div class="icon icon-largeMoney smallPTwo"></div><div class="text smallPTwo moneyPadding">Price</div></div></div>'+
                        ' <div class="modal-footer"><a  id="model-container-n" onclick="loginOrNulogin('+id+');return false;" role="button" data-toggle="modal" class="btn btn-lg playNow" >Play Now</a></div><ction> </div>';
                $("#body").html(b);
                $("#moneyPrice").html("×"+data.data.Price)
                hover();
            }
        }.bind(this),
        error: function(error){
            console.log(error);
        }
    });
}

//游戏 demo 页面登录没登录的问题
function loginOrNulogin(id){

   $("#game_idDemo").val(id);
   $.ajax({
        url: 'http://127.0.0.1:8083/checkplay?Game_id=' + id,
       async: false,
        type: 'GET',
        // data: chargeViewDatas,
        dataType: "json",
        success: function (result) {
            var data = result;
            if (data.code == "0") {
                $("#model-container-n").attr("href", "#modal-container-801879");
            } else if (data.code == "1") {
                $("#model-container-n").attr("href", "#modal-container-1");
                $("#userMoneyPrice").html("×" + data.data.Money);

            } else if (data.code == "2") {
               window.location.href=data.data;
            }

        }.bind(this),
        error: function (error) {
            console.log(error);
        }

    });

}
//进入个人中心页面
function personal(){
    $.ajax({
        url: 'http://127.0.0.1:8083/personal',
        type: 'GET',
        // data: chargeViewDatas,
        dataType: "json",
        success: function(result){
            var data = result;
            var chargeviewtb1 =[];
            var dataTwo = data.data.Game_list;
            var e='<div class="container"><div class="col-xs-8 col-sm-8 col-md-6 noPadding"><img src="../static/img/1.png" class="logoImg" /></div></div>';
            var a = '<div class="container bodyDiv"><header><h4>'+data.data.User_Name+'</h4><div class="icon icon-line"></div>'+
                    '<h5>The number of gold coins</h5><div class="moreButton"><a href="#" class="btn playNow" onclick="logout()" >Logout </a></div>'+
                    '<i class="icon icon-largeMoney smallPOne"></i><div class="smallPOne moneyFont moneySize moneyPadding">×<span>'+data.data.Money+
                    '</span></div></header><section class="paddingTop paddingBottom"><div class="col-sm-12 col-xs-12 noPadding paddingTop paddingBottom">'+
                    '<h4>Has Purchased The Game</h4><div class="icon icon-line paddingTop paddingBottom"></div></div><div class="col-sm-12 paddingBottomTwo"><div class="row row-top">';
                    if(dataTwo!=null) {
                var b = "";
                for (var j = 0; j < dataTwo.length; j++) {
                    if(b==""){
                      b='<div class="col-sm-2 col-xs-6 noPadding contentTwo">'+
                        '<img src="'+dataTwo[j].Image_url+'"  class="imgTwo" />'+
                        '<div class="txtTwo"><h6>'+dataTwo[j].Game_name+'</h6><a href="#" class="btn playNow" onclick="gameDemo('+dataTwo[j].Game_id+')">Play Now</a>'+
                        '</div></div>';

                    } else {
                        b=b+'<div class="col-sm-2 col-xs-6 noPadding contentTwo">'+
                            '<img src="'+dataTwo[j].Image_url+'"  class="imgTwo" />'+
                            '<div class="txtTwo"><h6>'+dataTwo[j].Game_name+'</h6><a href="#" class="btn playNow" onclick="gameDemo('+dataTwo[j].Game_id+')">Play Now</a>'+
                            '</div></div>';
                    }
                }
                }
                var d = "</div></div></section></div>";
                chargeviewtb1.push(e+a + b + d);
                $("#body").html(chargeviewtb1.join(""));
                hover();

        }.bind(this),
        error: function(error){
            console.log(error);
        }
    });
}
//首页ajax渲染
function index(){
    $.ajax({
        url: 'http://127.0.0.1:8083/gethomegame',
        type: 'GET',
        // data: chargeViewDatas,
        dataType: "json",
        success: function(result){
            var data = result;
            //console.log(data);
            var chargeviewtb1 =[];
            $.each(data.data, function(i, c) {
                var dataTwo=c.Game_list;
                var e= "'"+c.Tags_id+"'";
                var a= '<div class="container bodyDiv"><header><span class="h4" >'
                    + c.Tag_name+'</span><div class="moreButton"><a href="#" class="btn showAll" onclick="gameMore('+e+')" >Show All > </a></div>'+
                    '</header><div class="icon icon-line"></div><section><div class="col-sm-12"><div class="row row-top"><div class="col-sm-4 noPadding padding-top contentOne" >'+
                    '<img class="imgOne" src="'+ dataTwo[0].Image_url+'"/><div class="txt"><h6>'+ dataTwo[0].Game_name+'</h6><a href="#" class="btn playNow" onclick="gameDemo('+dataTwo[0].Game_id+')">Play Now</a>'+
                    '</div></div>';
                var b="";
                for(var j=1;j < 9;j=j+4) {
                    var jOne="",jTwo="",jThree="",jFour="";
                    if (b == "") {
                        if(dataTwo[j]!=undefined){
                            jOne="<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j].Image_url + "' /><div class='txtTwo'><h6>"+ dataTwo[j].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j].Game_id+")'>Play Now</a>"+
                                "</div></div>";
                        }
                        if(dataTwo[j + 1]!=undefined){
                            jTwo= "<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j + 1].Image_url + "' /><div class='txtTwo'><h6>"+ dataTwo[j+1].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j+1].Game_id+")'>Play Now</a>"+
                                "</div></div>" ;
                        }
                        if(dataTwo[j + 2]!=undefined){
                            jThree= "<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j + 2].Image_url + "' /><div class='txtTwo'><h6>"+ dataTwo[j+2].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j+2].Game_id+")'>Play Now</a>"+
                                "</div></div>";
                        }
                        if(dataTwo[j + 3]!=undefined){
                            jFour="<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j + 3].Image_url + "' /><div class='txtTwo'><h6>"+ dataTwo[j+3].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j+3].Game_id+")'>Play Now</a>"+
                                "</div></div>";
                        }
                        b="<div class='col-sm-4 noPadding'>"+jOne+jTwo+jThree+jFour+"</div>";
                    }else{
                        if(dataTwo[j]!=undefined){
                            jOne="<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j].Image_url + "' /><div class='txtTwo'><h6>"+ dataTwo[j].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j].Game_id+")'>Play Now</a>"+
                                "</div></div>";
                        }
                        if(dataTwo[j + 1]!=undefined){
                            jTwo= "<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j + 1].Image_url + "'/><div class='txtTwo'><h6>"+ dataTwo[j+1].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j+1].Game_id+")'>Play Now</a>"+
                                "</div></div>";
                        }
                        if(dataTwo[j + 2]!=undefined){
                            jThree= "<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j + 2].Image_url + "' /><div class='txtTwo'><h6>"+ dataTwo[j+2].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j+2].Game_id+")'>Play Now</a>"+
                                "</div></div>" ;
                        }
                        if(dataTwo[j + 3]!=undefined){
                            jFour="<div class='col-sm-6 col-xs-6 noPadding contentTwo'><img class='imgTwo' src='" + dataTwo[j + 3].Image_url + "'/><div class='txtTwo'><h6>"+ dataTwo[j+3].Game_name+"</h6><a href='#' class='btn playNow' onclick='gameDemo("+dataTwo[j+3].Game_id+")'>Play Now</a>"+
                                "</div></div>";
                        }
                        b = b+"<div class='col-sm-4 noPadding'>" +jOne+jTwo+jThree+jFour+ "</div>";

                    }
                }
                var d="</div></div></section></div>";
                chargeviewtb1.push(a+b+d);
            });
            $("#body").html(chargeviewtb1.join(""));
            hover();
        }.bind(this),
        error: function(error){
            console.log(error);
        }
    });

}
//用户是否登录
function gameLogin(){
    $.ajax({
        url: 'http://127.0.0.1:8083/login',
        type: 'GET',
        // data: chargeViewDatas,
        dataType: "json",
        success: function(result){
            var data=result;
            if(data.code=="0"){
              var a='<a id="modal-801879" href="#modal-container-801879" role="button" class="btn signButton" data-toggle="modal"><div class="icon icon-sign" ></div></a>';
                $("#sign").html(a);
            }else if(data.code=="1") {
                var dataTwo = data.data;
                var b = '<div class="signTwo"><a href="#" class="btn showAll smallScreen" onclick="personal()">' +
                    ' <i class="icon icon-user"></i><span id="user_name">' + dataTwo.User_Name + '</span><input type="hidden" id="user_email" value="' + dataTwo.Email + '" /></a>' +
                    '<div class="smallPTwo modal-bodyText line"><i class="icon icon-money"></i>' +
                    '<div class="smallPTwo paddingTop"><span class="moneyFont">×</span>' +
                    '<lable class="moneyFont">' + dataTwo.Money + '</lable></div></div></div>'
                $("#sign").html(b);
            }
        }.bind(this),
        error: function(error){
            console.log(error);
        }
    });
}
//退出登录
function logout(){
    $.ajax({
        url: 'http://127.0.0.1:8083/logout',
        type: 'GET',
        // data: chargeViewDatas,
        dataType: "json",
        success: function(result){
            var data = result;
         index();
         gameLogin();


        }.bind(this),
        error: function(error){
            console.log(error);
        }
    });
}
//游戏遮罩
function hover(){
    //游戏图片遮罩的js

    $(".contentOne").hover(function(){
        $(this).children(".txt").stop().animate({height:$(".imgOne").height()},200);
        $(this).find(".txt h6").stop().animate({paddingTop:$(".imgOne").height()*0.3},550);
        $(this).find(".txt .playNow").stop().show();
    },function(){
        $(this).children(".txt").stop().animate({height:"30px"},200);
        $(this).find(".txt h6").stop().animate({paddingTop:"0px"},550);
        $(this).find(".txt .playNow").stop().hide();
    });
    $(".contentTwo").hover(function(){
        $(this).children(".txtTwo").stop().animate({height:$(".imgTwo").height()},200);
        //$(this).children(".txtThree").stop().animate({height:$(".img-two").height()},200);
        $(this).find(".txtTwo h6").stop().animate({paddingTop:$(".imgTwo").height()*0.2},550);
        $(this).find(".txtTwo .playNow").stop().show();
    },function(){
        $(this).children(".txtTwo").stop().animate({height:"30px"},200);
        // $(this).children(".txtThree").stop().animate({height:"0px"},200);
        $(this).find(".txtTwo h6").stop().animate({paddingTop:"0px"},550);
        $(this).find(".txtTwo .playNow").stop().hide();
    });
}