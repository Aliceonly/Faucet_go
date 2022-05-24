function sorry_info() {
    var
        closeInSeconds = 5,
        displayText = ' #1 秒后将自动跳转发布订单页面',
        timer;
    swal({
        title: "抱歉,未找到您的订单!",
        text: displayText.replace(/#1/, closeInSeconds),
        imageUrl: "../static/image/info1.png",
        timer: closeInSeconds * 1000,
        showCancelButton: true,
    }, function () {
        window.location.href = "/post_job"
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
data = JSON.parse(sessionStorage.getItem('Query_hadAcceptOrder'))
console.log(data);
if (data == undefined) {
    var result = "";
    result += `<div class="row">
    <div class="col-lg-8">
    <div class="row">
    <div class="col-md-6 col-sm-6">
    <div class="blog-card">
    <div class="blog-text">
    <ul>
    <li>
    <i class='bx bxs-user'></i>
    <span>not found</span>
    </li>
    <li>
    <i class='bx bx-calendar'></i>
    <span>not found</span>
    </li>
    </ul>
    <h3>
    <a href="blog-details.html">not found</a>
    </h3>
    <p>金额：not found</p>
    <p>类型：not found</p>
    <p>时间戳：not found</p>
    </div>
    </div>
    </div>
    </div>
    </div>
     </div>`
    document.getElementById("release_order").innerHTML = result;
    sorry_info()

} else {
    var result = "";
    data.forEach(e => {
        // ${e.title}
        result += `<div class="row">
    <div class="col-lg-8">
    <div class="row">
    <div class="col-md-6 col-sm-6">
    <div class="blog-card">
    <div class="blog-text">
    <ul>
    <li>
    <i class='bx bxs-user'></i>
    <span>${e.State}</span>
    </li>
    <li>
    <i class='bx bx-calendar'></i>
    <span>${e.LaunchTime}</span>
    </li>
    </ul>
    <h3>
    <a>${e.Taskname}</a>
    </h3>
    <p>金额：${e.Amount}</p>
    <p>类型：${e.Category}</p>
    <p id="Timestamp_new">时间戳：${e.Timestamp}</p>
    <p id="Timestamp_new">hash:${e.Block}</p>
    <a onclick="comfirm_task(this)" class="blog-btn">
    <button>确认完成订单</button>
    <i class='bx bx-plus bx-spin'></i>
    </a>
    </div>
    </div>
    </div>
    </div>
    </div>
     </div>`
    })

    document.getElementById("release_order").innerHTML = result;
}

 
 

function comfirm_task(e) {
    var Timestamp_new = e.parentElement.children[4].innerText
    var hash_new = e.parentElement.children[5].innerText
    Timestamp = Timestamp_new.slice(4,14)
    hash = hash_new.slice(5,71)
    console.log(Timestamp_new);
    console.log(Timestamp);
    console.log(hash_new);
    console.log(hash);

    swal({
        title: "请输入您的签名",
        text: "确保您的签名无误，否则无法进行确认操作:",
        type: "input",
        showCancelButton: true,
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

    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/ClaimTrust",
        data: { timestap: Timestamp ,hash:hash,sign:inputValue},
        success: function (data){
            if(data.data=="ClaimOK"){
            console.log("success data", data);
            result = data.data
            console.log("result=====>", result);
            swal("Good!", "确认成功", "success");
        }else{
            swal("OMG", "删除操作失败了!", "error");
        }
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
   
});
}






























