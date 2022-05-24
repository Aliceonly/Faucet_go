function CreatTask() {
  var taskname = $("#exampleInput1").val();
  var tasktime = $("#exampleInput2").val();
  var tasktime2 = $("#exampleInput4").val();
  var taskmoney = $("#exampleInput3").val();
  var taskplace1 = $("#place option:selected").text();//地区
  var taskplace2 = $("#place2 option:selected").text();//主要地区
  var taskplace3 = $("#place3 option:selected").text();//工作类型
  var taskcontent = $('#exampleFormControlTextarea1').val();
  console.log(taskname, tasktime, taskmoney, tasktime2);
  console.log(taskplace1, taskplace2, taskplace3, taskcontent);
  var account=window.sessionStorage.getItem("Global_Account");
  $.ajax({
    method: "post",
    url: "http://localhost:8080/dapp/creatTask",
    data: { taskname: taskname, tasktime: tasktime + ":" + tasktime2, account: account, taskmoney: taskmoney, taskplace1: taskplace1 + taskplace2, taskplace3: taskplace3, taskcontent: taskcontent },
    beforeSend: function () {
      swal({
        title: "正在发布中，请稍等几秒.....",
        text:'<span style="color:red">请不要离开此页面、直至下一个弹框出现</br>否则交易可能失败!</sapn>',
        html:true,
        imageUrl: "../static/image/wait.png",
        showconfirmButton: true,
      })
    },
    success: function (data) {
      console.log("success data", data);
      console.log("成功");
      //  window.location.href ="http://localhost:8080/create_succ"
      swal({
        title: "发布成功",
        text: '您的订单时间戳是：<span style="color:red">' + data.data + '<br/>（时间戳可用于订单查询）</span><br/><a style="color:#3b3bf4" href="/"> 点击返回首页查看</a><br/>10秒后自动关闭。',
        imageUrl: "../static/image/check.png",
        html: true,
        timer: 5000,
        showConfirmButton: false
      });
    },
    error: function (data) {
      console.log("error====>", error)
      console.log("error data===>", data)
      swal("OMG!", "发布失败请重试！", "error");
    }
  })
}