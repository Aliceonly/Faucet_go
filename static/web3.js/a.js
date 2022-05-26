function demo_1() {
    swal("这是一个信息提示框!");
};
function demo_2() {
    swal("Good!", "登陆成功", "success");
};
function demo_3() {
    swal("OMG!", "登录失败，密码错误请重试！", "error");
};
function demo_4() {
    swal({
        title: "您确定要删除吗？",
        text: "您确定要删除这条数据？",
        type: "warning",
        showCancelButton: true,
        closeOnConfirm: false,
        confirmButtonText: "是的，我要删除",
        confirmButtonColor: "#ec6c62"
    }, function () {
        $.ajax({
            url: "do.php",
            type: "DELETE"
        }).done(function (data) {
            swal("操作成功!", "已成功删除数据！", "success");
        }).error(function (data) {
            swal("OMG", "删除操作失败了!", "error");
        });
    });
};
function demo_5() {
    swal({
        title: "发布成功",
        text: '<span style="color:red">您已发布成功，</span><a style="color:#3b3bf4" href="inde.html">点击返回首页查看</a>。<br/>5秒后自动关闭。',
        imageUrl: "check.png",
        html: true,
        timer: 5000,
        showConfirmButton: false
    });
};
function demo_6() {
    swal({
        title: "请输入您的签名",
        text: "确保您的签名无误，否则无法进行确认操作:",
        type: "input",
        showCancelButton: true,
        confirmButtonColor:"#3085d6",
        confirmButtonColor:true,
        closeOnConfirm: false,
        animation: "slide-from-top",
        inputPlaceholder: "签名"
    }, function (inputValue) {
        if (inputValue === false) return false;
        if (inputValue === "") {
            swal.showInputError("请输入!");
            return false
        }
       
        console.log("test======>",inputValue);
          swal("Good!", "确认成功", "success");
    });
    
};
 
 

function demo71() {
    var
        closeInSeconds = 5,
        displayText = ' #1 秒后将自动跳转登录页面',
        timer;

    swal({
        title: "请进行登录!",
        text: displayText.replace(/#1/, closeInSeconds),
        imageUrl: "../static/image/info1.png",
        timer: closeInSeconds * 1000,
        showCancelButton: true, //有这个就有取消按钮
        showconfirmButton: true,
    }, function () {
            window.location.href = "/account"
        }
    );

    timer = setInterval(function () {
        closeInSeconds--;
        if (closeInSeconds < 0) {
            clearInterval(timer);
        }

        $('.sweet-alert > p').text(displayText.replace(/#1/, closeInSeconds));

    }, 1000);
}

function demo7(){
    swal({
        title: "正在发布中，请稍等几秒.....",
        text:'<span style="color:red">请不要离开此页面、直至下一个弹框出现</br>否则交易可能失败!</sapn>',
        html:true,
        imageUrl: "../static/image/wait.png",
        showconfirmButton: true,
      })
}
// swal({
    // title: "您确定要删除吗？",
    // text: "您确定要删除这条数据？",
    // type: "warning",
    // showCancelButton: true,
    // closeOnConfirm: false,
    // confirmButtonText: "是的，我要删除",
    // confirmButtonColor: "#ec6c62"
// }, function () {
    // $.ajax({
        // url: "do.php",
        // type: "DELETE"
    // }).done(function (data) {
        // swal("操作成功!", "已成功删除数据！", "success");
    // }).error(function (data) {
        // swal("OMG", "删除操作失败了!", "error");
    // });
// });
