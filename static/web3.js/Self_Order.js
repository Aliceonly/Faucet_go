let Account = window.sessionStorage.getItem("Global_Account")
console.log(Account)
document.getElementById("Self_account").innerHTML = Account;
function release_order() {
    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/Self_Order_show",
        data: { Account: Account},
        success: function (data) {
            sessionStorage.setItem('Query_release_order',JSON.stringify(data.data))
            console.log(data.data)
            window.location.href="/release_order"
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}