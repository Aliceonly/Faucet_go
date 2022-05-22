const req = require("express/lib/request");
const res = require("express/lib/response");


function create_user(){
    var a=1,b=1,c=1,d=1,e=1,f=1,g=1;
    var Sid=$("#Sid").val();
    var Telephone=$("#Telephone").val();
    var Password=$("#password").val();
    var CPassword=$("#passwords").val();
    // var phonelength = $("#Telephone").val().length;
    reg = /^1(3\d|4\d|5\d|6\d|7\d|8\d|9\d)\d{8}$/g 
    var result = reg.test(Telephone);
    console.log(result)
    if(result == false){
        swal("Sorry!", "请输入正确的手机号码！", "error");
        a=0
        console.log(a)
    }else{
        a=1
    }
    if (Telephone == "") {
        swal("Sorry!", "手机号码为空，请重试！", "error");
        // alert("请输入手机号码！");
        // return false;
        b=0
    }else{
        b=1
    } 
    if (Sid == "") {
        swal("Sorry!", "学号为空，请重试！", "error");
        // alert("请输入学号！");
        // return false;
        c=0
    } else{
        c=1
    }
    if (Password == "") {
        swal("Sorry!", "密码为空，请重试！", "error");
        // alert("请输入密码！");
        // return false;
        d=0
    } else{
        d=1
    }
    if (Password == ""&&Sid == ""&&Telephone == "") {
        swal("Sorry!", "输入为空，请重试！", "error");
        e=0
    }else{
        e=1
    }
    if (CPassword == Password) {
        f=1
    }else{
        swal("Sorry!", "密码不一致，请重试！", "error");
        f=0
    }
    console.log(Sid,Telephone,Password,CPassword)
    if (a == 1 & b == 1 & c == 1 & d == 1 & f == 1){
        $.ajax({
            method:"post",
            url: "http://localhost:8080/dapp/creatUser",
            data: { Sid:Sid, Telephone:Telephone, Password: Password },
            success: function (data) {
                // swal("注册成功!", "您的账户是: " + data.data, "success");
                swal({
                    title: "注册成功",
                    text: '<span>您的账户是：</span><br/><span style="color:red">'+ data.data +'</span><br/><span style="color:red">请务必记住，一旦丢失无法找回<span><br/><a style="color:#3b3bf4" href="/sign_in">点击进行登录</a><br/>10秒后自动关闭。',
                    imageUrl: "../static/image/check.png",
                    html: true,
                    timer: 10000,
                    showConfirmButton: false
                });
                console.log("success insert user",data);
            },
            error: function (data) {
                console.log("error====>", error)
                console.log("error data===>", data)
            }
        })
    }else {
        console.log("errrr")
    }
}

 