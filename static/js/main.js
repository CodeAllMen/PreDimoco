// $(document).ready(function() {
//     // 网页初加载时页面展示
//     $("#modal-container-1").modal("show");
//   });
  
  
  new Fingerprint2().get(function(result, components) {
    console.log(result); // a hash, representing your device fingerprint
    // console.log(components) // an array of FP components
    $.ajax({
      type: "GET",
      url: "http://127.0.0.1:8095/returnid",
      data: {
        type: "video_w",
        affName: GetQueryString("affName"),
        clickId: GetQueryString("clickId"),
        pubId: GetQueryString("pubId"),
        proId: GetQueryString("proId"),
        canvasId: result,
        ct: "NL"
      }
    }).done(function(result) {
      ptxid = result;
      if (ptxid == "false") {
        window.location.href = "http://google.com";
      }
    });
  });
  
  function closeWin() {
    if (
      navigator.userAgent.indexOf("Firefox") != -1 ||
      navigator.userAgent.indexOf("Chrome") != -1
    ) {
      window.location.href = "about:blank";
      window.close();
    } else {
      window.opener = null;
      window.open("", "_self");
      window.close();
    }
  }
  
  function GetQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;
  }
  
  
  
  $("#submit_btn_mobile").click(function() {
    window.location.href = "http://cpx5.allcpx.com/subscript/request/" + ptxid;
  });
  
  $("#next").click(function() {
    $("#grayheader").removeClass("total");
    $("#mobile_box").removeClass("total");
    $("#modal-container-1").modal("hide");
  });
  
  $("#back").click(function() {
    window.location.href = "http://google.com";
  });
  