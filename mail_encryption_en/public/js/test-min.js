window.onload = function() { var mailList = document.getElementById("mailList"); if (mailList != null) { mailListUp() } var setma = document.getElementById("setMail"); if (setma != null) setma.addEventListener('DOMContentLoaded', setmailListUp()); var setse = document.getElementById("setMailSer"); if (setse != null) setse.addEventListener('DOMContentLoaded', setMailServer()); var add = document.getElementById("set_add"); if (add != null) { add.addEventListener("click", function(e) { var plusUl = document.createElement('option');
            plusUl.innerHTML = '새로운 이메일</option>';
            document.getElementById('mailList').appendChild(plusUl) }, false) } var logout = document.getElementById("logout"); if (logout != null) { logout.addEventListener("click", function(e) { console.log("로그아웃 실행") }, false);
        logout.addEventListener("click", logoutClick, false) } var setBtn = document.getElementById("set"); if (setBtn != null) { var xhr = new XMLHttpRequest();
        xhr.onload = function() { if (xhr.status === 200 || xhr.status === 201) { var setAdd = document.getElementById("setAdd"); var setDel = document.getElementById("setDel"); var setSave = document.getElementById("setSave"); var setId = document.getElementById("setId"); var setSelect = document.getElementById("menuMailList");
                menvaruSelect = document.getElementById("set_mailList"); if (setAdd != null && setId != null) { var add = document.getElementById("set_add"); if (add != null) { add.addEventListener("click", email, false) } }
                search = xhr.responseText; var jbSplit = search.split(';');
                jbSplit.length = jbSplit.length - 1; for (var i in jbSplit) { var addOp = document.createElement('option');
                    addOp.innerHTML = jbSplit[i] + '</option>';
                    document.getElementById('set_mailList').appendChild(addOp);
                    console.log(i + ": " + jbSplit[i]); if (i + 1 == jbSplit.length) break } } else { alert("오류 발생");
                window.location.reload() } };
        xhr.open('GET', '/set', true);
        xhr.send(null) } };

function post(URL, PARAMS) { var temp = document.createElement("form");
    temp.action = URL;
    temp.method = "post";
    temp.style.display = "none"; for (var x in PARAMS) { var opt = document.createElement("textarea");
        opt.name = x;
        opt.value = PARAMS[x];
        temp.appendChild(opt) }
    document.body.appendChild(temp);
    temp.submit() }

function page_move(s_page) { post("/compose", { bo: s_page, }) }

function Write() { var to = document.getElementById("writeTo").value; var title = document.getElementById("enTitle").value; var input = document.getElementById("enSource").value; var xhr = new XMLHttpRequest(); var data = { jto: to, jtitle: title, jinput: input, };
    xhr.onload = function() { if (xhr.status === 200 || xhr.status === 201) { alert("전송 완료!");
            window.history.back() } else { console.error(xhr.responseText);
            alert("오류 발생");
            window.location.reload() } };
    xhr.open('POST', '/write');
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify(data)) }

function goBack() { window.history.back() }

function setEmailServer(e) { var chk = document.querySelector('input[name="chk_info"]:checked').value;
    console.log(chk);
    console.log(e); var ia = document.getElementById("imap_add"); var ip = document.getElementById("imap_port"); var sa = document.getElementById("smtp_add"); var sp = document.getElementById("smtp_port"); var ma = document.getElementById("mail"); var pw = document.getElementById("mail_passwd"); var xhr = new XMLHttpRequest();
    xhr.onload = function() { if (xhr.status === 200 || xhr.status === 201) { search1 = xhr.response;
            console.log("1" + search1); var obj = JSON.parse(search1);
            console.log("2" + obj.imap_add);
            ia.value = obj.imap_add;
            ip.value = obj.imap_port;
            sa.value = obj.smtp_add;
            sp.value = obj.smtp_port } else { console.error(xhr.responseText);
            alert("오류 발생");
            window.location.reload() } };
    xhr.open("GET", "/modMailServer?chk=" + chk + "&mail=" + e, true);
    xhr.send() }

function mailListUp(e) { var xhr = new XMLHttpRequest();
    xhr.onload = function() { if (xhr.status === 200 || xhr.status === 201) { search = xhr.responseText; var jbSplit = search.split(';');
            jbSplit.length = jbSplit.length - 1;
            console.log("이메일 리스트: " + search); for (var i in jbSplit) { var addOp = document.createElement('option');
                addOp.innerHTML = jbSplit[i] + '</option>';
                document.getElementById('mailList').appendChild(addOp);
                console.log(i + ": " + jbSplit[i]); if (i + 1 == jbSplit.length) break }
            console.log("mailListVal in Up: " + sessionStorage.getItem('eid')); if (sessionStorage.getItem('eid') != null) { for (i = 0; i < document.getElementById("mailList").options.length; i++) { if (document.getElementById("mailList").options[i].value == sessionStorage.getItem('eid')) { document.getElementById("mailList").options[i].selected = "selected" } } } } else { alert("오류 발생");
            window.location.reload() } };
    xhr.open('POST', '/mailList', true);
    xhr.setRequestHeader('Content-Type', 'text/plain');
    xhr.send() }

function setmailListUp(e) { var xhr = new XMLHttpRequest();
    xhr.onload = function() { if (xhr.status === 200 || xhr.status === 201) { search = xhr.responseText; var jbSplit = search.split(';');
            jbSplit.length = jbSplit.length - 1;
            console.log("설정 이메일 서버1" + search); for (var i in jbSplit) { var addOp = document.createElement('option');
                addOp.setAttribute("value", jbSplit[i]);
                addOp.innerHTML = jbSplit[i] + '</option>';
                document.getElementById('setMailList').appendChild(addOp); if (i + 1 == jbSplit.length) break } } else { alert("오류 발생");
            window.location.reload() } };
    xhr.open('POST', '/mailList', true);
    xhr.setRequestHeader('Content-Type', 'text/plain');
    xhr.send() }

function setMailServer(e) { var xhr = new XMLHttpRequest();
    xhr.onload = function() { if (xhr.status === 200 || xhr.status === 201) { search = xhr.responseText;
            console.log("설정 이메일 서버2" + search); var jbSplit = search.split(';');
            jbSplit.length = jbSplit.length - 1; for (var i in jbSplit) { var addOp = document.createElement('option');
                addOp.setAttribute("value", jbSplit[i]);
                addOp.innerHTML = jbSplit[i] + '</option>';
                document.getElementById('setMailServer').appendChild(addOp);
                console.log("설정 이메일 서버" + i + ": " + jbSplit[i]); if (i + 1 == jbSplit.length) break } } else { alert("오류 발생");
            window.location.reload() } };
    xhr.open('POST', '/mailPreset', true);
    xhr.setRequestHeader('Content-Type', 'text/plain');
    xhr.send() }

function mailList_ch(e) { var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() { if (this.readyState == 4 && this.status == 200) { search = xhr.responseText; var jbSplit = search.split(';');
            jbSplit.length = jbSplit.length - 1; for (var i in jbSplit) { var addOp = document.createElement('option');
                addOp.innerHTML = jbSplit[i] + '</option>';
                document.getElementById('mailList').appendChild(addOp);
                console.log(i + ": " + jbSplit[i]); if (i + 1 == jbSplit.length) break }
            sessionStorage.setItem('eid', e);
            console.log("eid in Ch: " + sessionStorage.getItem('eid')); if (search == "1") { location.replace("/index") } else { location.reload() } } }; var para = document.location.href.split("/");
    console.log(para);
    xhr.open("GET", "/mailChange?id=" + e + "&url=" + para[3], true);
    xhr.send() }

function setListClick(e) { var ia = document.getElementById("imap_add"); var ip = document.getElementById("imap_port"); var sa = document.getElementById("smtp_add"); var sp = document.getElementById("smtp_port"); var ma = document.getElementById("mail"); var pw = document.getElementById("mail_passwd"); var tempmail = document.getElementById("tempmail"); var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() { if (this.readyState == 4 && this.status == 200) { search1 = xhr.response;
            console.log(search1); var obj = JSON.parse(search1);
            console.log(obj.imap_add);
            ia.value = obj.imap_add;
            ip.value = obj.imap_port;
            sa.value = obj.smtp_add;
            sp.value = obj.smtp_port;
            ma.value = obj.mail;
            pw.value = obj.mail_passwd;
            tempmail.value = obj.mail } };
    xhr.open("GET", "/modMailServer?mail=" + e, true);
    xhr.send() }

function setAdd(e) { var iAdd = document.getElementById("imap_add").value; var iPort = document.getElementById("imap_port").value; var sAdd = document.getElementById("smtp_add").value; var sPort = document.getElementById("smtp_port").value; var mail = document.getElementById("mail").value; var mailPw = document.getElementById("mail_passwd").value; if (iAdd != "" || iPort != "" || sAdd != "" || sPort != "" || mail != "" || mailPw != "") { var fm = document.getElementById("setf");
        fm.method = "post";
        fm.target = "_self";
        fm.action = "/mailserverInsert";
        fm.submit() } else { alert("모든 항목을 입력하세요.") } }

function setDel(e) { var fm = document.getElementById("setfs"); var email = $("#setMailList option:selected").val();
    post("/mailserverDelete", { mail: email, }) }

function setMod(e) { var iAdd = document.getElementById("imap_add").value; var iPort = document.getElementById("imap_port").value; var sAdd = document.getElementById("smtp_add").value; var sPort = document.getElementById("smtp_port").value; var mail = document.getElementById("mail").value; var mailPw = document.getElementById("mail_passwd").value; var tempmail = document.getElementById("tempmail"); if (iAdd != "" || iPort != "" || sAdd != "" || sPort != "" || mail != "" || mailPw != "" || tempmail != "") { var fm = document.getElementById("setf");
        fm.method = "post";
        fm.target = "_self";
        fm.action = "/mailserverUpdate";
        fm.submit() } else { alert("모든 항목을 입력하세요.") } }

function setDefault(e) { var mail = document.getElementById("setMailList"); if (mail.value != "") { mail = mail.options[mail.selectedIndex].value; var opt = document.createElement("input");
        opt.value = mail;
        opt.type = "hidden";
        opt.name = "opt"; var fm = document.getElementById("setf");
        fm.appendChild(opt);
        fm.method = "post";
        fm.target = "_self";
        fm.action = "/defaultmailChange";
        fm.submit() } else { alert("이메일을 선택하세요.") } }

function logoutClick(e) { alert("로그인 페이지로 이동합니다.") }

function write_enck() { var enck = document.getElementById("enck"); var en_sub = document.getElementById("en_sub"); if (enck.checked == true) { document.getElementById("en_sub").style.display = "block" } else document.getElementById("en_sub").style.display = "none" }