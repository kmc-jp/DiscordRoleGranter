const encodeURIParams = (param) => {
    let url = "";
    for (p of Object.keys(param)) {
        url += p + "=" + encodeURIComponent(param[p]) + "&";
    }
    return url.slice(0, -1);;
}

const url_make = (param) => {
    return location.protocol + "//" + location.hostname + location.pathname + "?" + encodeURIParams(param);
}

const make_alert = (text, mode) => {
    let alertdiv = document.getElementById("alert");
    alertdiv.innerHTML = ""
    switch (String(mode)) {
        case "error":
            alertdiv.className = "alert alert-danger";
            alertdiv.innerText = text;
            break;
        case "info":
            alertdiv.className = "alert alert-info";
            alertdiv.innerText = text;
            break;
        case "remove":
            alertdiv.className = "";
            break;
        default:
            alertdiv.className = "alert alert-success";
            alertdiv.innerText = text;
    }
}

const find_member = async() => {

    let res = await fetch(url_make({ "Action": "List" }), {
        method: "Get",
        credentials: 'same-origin',
    })

    let result = await res.json()
    if (res.status !== 200) {
        make_alert("ユーザ一覧の取得に失敗しました。(" + result.Status + ")", "error")
        return
    }

    let value = document.querySelector("#find_content").value.toLowerCase().trim();
    if (value === "") {
        return
    }

    let list = []
    result.Member.forEach(mem => {
        if (mem.UserName.toLowerCase().includes(value)) {
            list.push(mem)
            return
        }
        if (mem.Nick.toLowerCase().includes(value)) {
            list.push(mem)
            return
        }
        if (mem.ID === value) {
            list.push(mem)
            return
        }
    })

    let table = ""
    list.forEach(mem => {
        table +=
            '<tr onclick="add_role(\'' + mem.ID + '\')">' +
            '<td>' + mem.UserName + '</td>' +
            '<td>' + mem.Nick + '</td>' +
            '<td>' + mem.ID + '</td>' +
            '</tr>'
    })

    if (list.length < 1) {
        document.querySelector("#find_res").textContent = "該当するユーザが見つかりませんでした。"
        return
    }

    document.querySelector("#find_res").innerHTML =
        '<p>以下から目的のユーザを選択してください。</p>' +
        '<table class="table">' +
        '<thead>' +
        '<tr>' +
        '<th scope="col">UserName</th>' +
        '<th scope="col">NickName</th>' +
        '<th scope="col">ID</th>' +
        '</tr>' +
        '</thead>' +
        '<tbody>' +
        table +
        '</tbody > '
}

const add_role = async(user_id) => {
    let res = await fetch(url_make({ "Action": "AddRole" }), {
        method: "POST",
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        credentials: 'same-origin',
        body: encodeURIParams({ "user_id": user_id })
    })

    let body = await res.json()

    if (res.status !== 200) {
        make_alert("失敗しました。(" + body["Status"] + ")", "error")
        return
    }

    make_alert("成功しました。")
    document.querySelector("#user_id").value = ""
    return

}