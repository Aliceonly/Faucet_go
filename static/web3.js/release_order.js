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
data = JSON.parse(sessionStorage.getItem('Query_release_order'))
console.log(data);
if(data == undefined){
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

}else{
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
    <a href="blog-details.html">${e.Taskname}</a>
    </h3>
    <p>金额：${e.Amount}</p>
    <p>类型：${e.Category}</p>
    <p>时间戳：${e.Timestamp}</p>
    <a href="blog-details.html" class="blog-btn">
    Read More
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



























