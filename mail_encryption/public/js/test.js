window.onload = function() {
    var mailList = document.getElementById("mailList");
    if (mailList != null) {
        mailListUp();
    }
    // setmailListUp();
    // setMailServer();

    var setma = document.getElementById("setMail");
    if (setma != null)
        setma.addEventListener('DOMContentLoaded', setmailListUp());
    var setse = document.getElementById("setMailSer");
    if (setse != null)
        setse.addEventListener('DOMContentLoaded', setMailServer());

    var add = document.getElementById("set_add");
    if (add != null) {
        add.addEventListener("click", function(e) {
            var plusUl = document.createElement('option');
            plusUl.innerHTML = '새로운 이메일</option>';
            document.getElementById('mailList').appendChild(plusUl);
        }, false);
    }
    var logout = document.getElementById("logout");
    if (logout != null) {
        logout.addEventListener("click", function(e) {
            console.log("로그아웃 실행");
        }, false);
        logout.addEventListener("click", logoutClick, false);
    }

    //설정 - 추가, 삭제, 저장
    var setBtn = document.getElementById("set");
    if (setBtn != null) {

        var xhr = new XMLHttpRequest();

        //    console.log("in js pw: " + pw);
        xhr.onload = function() { // get /set 완료후 페이지 로드가 되면 이메일서버 선택 리스트 / 하단 이메일 리스트를 리스트업, 체크박스와 이메일 선택 리스트 클릭 시 db로부터 하단 2줄 채우기
            if (xhr.status === 200 || xhr.status === 201) {
                // setEmailServer();
                var setAdd = document.getElementById("setAdd");
                var setDel = document.getElementById("setDel");
                var setSave = document.getElementById("setSave");

                var setId = document.getElementById("setId");
                var setSelect = document.getElementById("menuMailList")
                var menuSelect = document.getElementById("set_mailList")
                    // $("#SelectBoxID").change(setEmailServer());

                if (setAdd != null && setId != null) {
                    var add = document.getElementById("set_add");
                    if (add != null) {
                        add.addEventListener("click", email, false);
                    }
                }

                //console.log(xhr.responseText);
                search = xhr.responseText;
                var jbSplit = search.split(';');
                jbSplit.length = jbSplit.length - 1;

                for (var i in jbSplit) {
                    var addOp = document.createElement('option');
                    addOp.innerHTML = jbSplit[i] + '</option>';
                    document.getElementById('set_mailList').appendChild(addOp);
                    console.log(i + ": " + jbSplit[i]);
                    if (i + 1 == jbSplit.length) break;
                }
                // alert("전송 완료!");
            } else {
                //console.error(xhr.responseText);
                alert("오류 발생");
                window.location.reload();
            }
        };
        xhr.open('GET', '/set', true);
        // xhr.setRequestHeader('Content-Type', 'text/plain');
        xhr.send(null);
    }


    // fetch('menuEmailList').then(function(response) {
    //     response.text().then(function(text) {
    //         console.log(text);
    //         var items = text.split(',');
    //         console.log(items);
    //         var i = 0;
    //         var tags = '';
    //         while (i < items.length) {
    //             var item = items[i];
    //             //<option value="gmail">a123@gmail.com</option>
    //             var tag = '<option>' + item + '</option>'
    //             tags += tag;
    //             i += 1;
    //         }

    //         document.querySelector("#mailList").innerHTML = tags;
    //     })
    // });

};

// 단순 submit
function post(URL, PARAMS) {
    var temp = document.createElement("form");
    temp.action = URL;
    temp.method = "post";
    temp.style.display = "none";
    for (var x in PARAMS) {
        var opt = document.createElement("textarea");
        opt.name = x;
        opt.value = PARAMS[x];
        temp.appendChild(opt);
    }
    document.body.appendChild(temp);
    temp.submit();
    //return temp;
}

//메일 목록에서 보낸 사람 클릭 시 메일쓰기로 이동
function page_move(s_page) {
    post("/compose", {
            bo: s_page,
        })
        // var temp = document.createElement("form");
        // var bo = document.createElement("input");
        // bo.value = s_page
        // temp.appendChild(bo)
        // temp.action = "/compose";
        // temp.method = "post";
        // temp.submit();
}



// 메일 보내기
function Write() {
    var to = document.getElementById("writeTo").value;
    // var cc = document.getElementById("write_cc").value;
    var title = document.getElementById("enTitle").value;
    var input = document.getElementById("enSource").value;
    var xhr = new XMLHttpRequest();
    var data = {
        jto: to,
        //jcc: cc,
        jtitle: title,
        jinput: input,
    };
    xhr.onload = function() {
        if (xhr.status === 200 || xhr.status === 201) {
            //console.log(xhr.responseText);
            alert("전송 완료!");
            window.history.back();
        } else {
            console.error(xhr.responseText);
            alert("오류 발생");
            window.location.reload();
        }
    };
    xhr.open('POST', '/write');
    xhr.setRequestHeader('Content-Type', 'application/json'); // 컨텐츠타입을 json으로
    //
    xhr.send(JSON.stringify(data)); // 데이터를 stringify해서 보냄
    // xhr.abort(); // 전송된 요청 취소
}

//로그인
// function loginPost() {
//     sessionStorage.setItem('eid', "");
//     console.log("on load: " + sessionStorage.getItem('eid'));
//     // var fm = document.getElementById("f1");
//     // fm.method = "post";
//     // fm.target = "_self";
//     // fm.action = "/loginTest";
//     // fm.submit();

//     //     // var xhr = new XMLHttpRequest();
//     //     // var id = document.getElementById("id").value;
//     //     // var pw = document.getElementById("pw").value;
//     //     // //    console.log("in js pw: " + pw);
//     //     // xhr.onload = function() {
//     //     //     if (xhr.status === 200 || xhr.status === 201) {
//     //     //         //console.log(xhr.responseText);
//     //     //         //   alert("전송 완료!");
//     //     //     } else {
//     //     //         //console.error(xhr.responseText);
//     //     //         alert("오류 발생");
//     //     //         window.location.reload();
//     //     //     }
//     //     // };
//     //     // xhr.open('POST', '/loginTest', true);
//     //     // xhr.setRequestHeader('Content-Type', 'text/plain');
//     //     // xhr.send(id + "," + pw);
//     //     // //  alert(xhr.response);
//     //     // // xhr.abort(); // 전송된 요청 취소
// }


// 이전 페이지 돌아가기
function goBack() {
    window.history.back();
}



// 설정 - 이메일 서버 셀렉트 선택 시 주소와 포트번호 자동 설정
function setEmailServer(e) {

    // $('#setMailServer option').click(function() {
    //     alert($(this).val());
    // });

    // var s = document.getElementById("selEmailServer").value;
    // var chk = document.getElementById("chk_info").value;
    var chk = document.querySelector('input[name="chk_info"]:checked').value;
    console.log(chk); //imap, pop
    console.log(e); // 메일 서버

    var ia = document.getElementById("imap_add");
    var ip = document.getElementById("imap_port");
    var sa = document.getElementById("smtp_add");
    var sp = document.getElementById("smtp_port");
    var ma = document.getElementById("mail");
    var pw = document.getElementById("mail_passwd");

    var xhr = new XMLHttpRequest();

    xhr.onload = function() {
        if (xhr.status === 200 || xhr.status === 201) {
            search1 = xhr.response;
            console.log("1" + search1);
            var obj = JSON.parse(search1);
            console.log("2" + obj.imap_add);

            // console.log(xhr.responseText);
            // search1 = xhr.response;
            // console.log(search);
            // var obj = JSON.parse(search1);

            ia.value = obj.imap_add;
            ip.value = obj.imap_port;
            sa.value = obj.smtp_add;
            sp.value = obj.smtp_port;
            // ma.value = "";
            // pw.value = "";
            // window.history.back();
        } else {
            console.error(xhr.responseText);
            alert("오류 발생");
            window.location.reload();
        }
    };

    xhr.open("GET", "/modMailServer?chk=" + chk + "&mail=" + e, true);
    xhr.send();
}


//메뉴 - 이메일 리스트 가져오기
function mailListUp(e) {
    var xhr = new XMLHttpRequest();
    //alert(sessionStorage.getItem('eid'));
    //    console.log("in js pw: " + pw);
    xhr.onload = function() {
        if (xhr.status === 200 || xhr.status === 201) {
            //console.log(xhr.responseText);
            search = xhr.responseText;
            var jbSplit = search.split(';');
            jbSplit.length = jbSplit.length - 1;
            //  $('#mailList').empty();
            console.log("이메일 리스트: " + search);
            for (var i in jbSplit) {
                var addOp = document.createElement('option');
                // addOp.setAttribute("value", i + 1);
                addOp.innerHTML = jbSplit[i] + '</option>';
                document.getElementById('mailList').appendChild(addOp);
                console.log(i + ": " + jbSplit[i]);
                if (i + 1 == jbSplit.length) break;
            }
            console.log("mailListVal in Up: " + sessionStorage.getItem('eid'));
            if (sessionStorage.getItem('eid') != null) {
                for (i = 0; i < document.getElementById("mailList").options.length; i++) {
                    if (document.getElementById("mailList").options[i].value == sessionStorage.getItem('eid')) {
                        document.getElementById("mailList").options[i].selected = "selected";
                    }
                }
            }
            // alert("전송 완료!");
        } else {
            //console.error(xhr.responseText);
            alert("오류 발생");
            window.location.reload();
        }
    };
    xhr.open('POST', '/mailList', true);
    xhr.setRequestHeader('Content-Type', 'text/plain');
    xhr.send();
}

//설정 - 이메일 리스트 가져오기
function setmailListUp(e) {
    var xhr = new XMLHttpRequest();

    // console.log("in js pw: " + pw);
    xhr.onload = function() {
        if (xhr.status === 200 || xhr.status === 201) {
            //console.log(xhr.responseText);
            search = xhr.responseText;
            var jbSplit = search.split(';');
            jbSplit.length = jbSplit.length - 1;
            //  $('#mailList').empty();
            console.log("설정 이메일 서버1" + search);
            for (var i in jbSplit) {
                var addOp = document.createElement('option');
                addOp.setAttribute("value", jbSplit[i]);
                addOp.innerHTML = jbSplit[i] + '</option>';
                document.getElementById('setMailList').appendChild(addOp);
                //console.log(i + ": " + jbSplit[i]);
                if (i + 1 == jbSplit.length) break;
            }
            // alert("전송 완료!");
        } else {
            //console.error(xhr.responseText);
            alert("오류 발생");
            window.location.reload();
        }
    };
    xhr.open('POST', '/mailList', true);
    xhr.setRequestHeader('Content-Type', 'text/plain');
    xhr.send();
}

//설정 - 이메일 서버 셀렉트 리스트 가져오기
function setMailServer(e) {
    var xhr = new XMLHttpRequest();

    //    console.log("in js pw: " + pw);
    xhr.onload = function() {
        if (xhr.status === 200 || xhr.status === 201) {
            //console.log(xhr.responseText);
            search = xhr.responseText;
            console.log("설정 이메일 서버2" + search);
            var jbSplit = search.split(';');
            jbSplit.length = jbSplit.length - 1;
            //  $('#mailList').empty();

            for (var i in jbSplit) {
                var addOp = document.createElement('option');
                addOp.setAttribute("value", jbSplit[i]);
                addOp.innerHTML = jbSplit[i] + '</option>';
                document.getElementById('setMailServer').appendChild(addOp);
                console.log("설정 이메일 서버" + i + ": " + jbSplit[i]);
                if (i + 1 == jbSplit.length) break;
            }
            // alert("전송 완료!");
        } else {
            //console.error(xhr.responseText);
            alert("오류 발생");
            window.location.reload();
        }
    };
    xhr.open('POST', '/mailPreset', true);
    xhr.setRequestHeader('Content-Type', 'text/plain');
    xhr.send();
}

//메뉴 메일리스트 선택
function mailList_ch(e) {

    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            //console.log(xhr.responseText);
            search = xhr.responseText;


            //xhr.response
            var jbSplit = search.split(';');
            jbSplit.length = jbSplit.length - 1;
            //  $('#mailList').empty();

            for (var i in jbSplit) {
                var addOp = document.createElement('option');
                addOp.innerHTML = jbSplit[i] + '</option>';
                document.getElementById('mailList').appendChild(addOp);
                console.log(i + ": " + jbSplit[i]);
                if (i + 1 == jbSplit.length) break;
            }
            sessionStorage.setItem('eid', e);
            console.log("eid in Ch: " + sessionStorage.getItem('eid'));
            if (search == "1") {
                location.replace("/index");
            } else {
                location.reload();
            }

        }
    };
    var para = document.location.href.split("/");
    console.log(para);

    xhr.open("GET", "/mailChange?id=" + e + "&url=" + para[3], true);
    xhr.send();

}

// 설정 - 메일 리스트 클릭
function setListClick(e) {
    //console.log("this.value e: " + e)
    var ia = document.getElementById("imap_add");
    var ip = document.getElementById("imap_port");
    var sa = document.getElementById("smtp_add");
    var sp = document.getElementById("smtp_port");
    var ma = document.getElementById("mail");
    var pw = document.getElementById("mail_passwd");
    var tempmail = document.getElementById("tempmail");

    var xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            //console.log(xhr.responseText);

            search1 = xhr.response;
            console.log(search1);
            var obj = JSON.parse(search1);
            console.log(obj.imap_add);

            ia.value = obj.imap_add;
            ip.value = obj.imap_port;
            sa.value = obj.smtp_add;
            sp.value = obj.smtp_port;
            ma.value = obj.mail;
            pw.value = obj.mail_passwd;
            tempmail.value = obj.mail;
            // sessionStorage.setItem('tempmail', obj.mail);
        }
    };

    xhr.open("GET", "/modMailServer?mail=" + e, true);
    xhr.send();
}

// 설정 이메일 추가
function setAdd(e) {

    var iAdd = document.getElementById("imap_add").value;
    var iPort = document.getElementById("imap_port").value;
    var sAdd = document.getElementById("smtp_add").value;
    var sPort = document.getElementById("smtp_port").value;
    var mail = document.getElementById("mail").value;
    var mailPw = document.getElementById("mail_passwd").value;

    if (iAdd != "" || iPort != "" || sAdd != "" || sPort != "" || mail != "" || mailPw != "") {
        var fm = document.getElementById("setf");
        fm.method = "post";
        fm.target = "_self";
        fm.action = "/mailserverInsert";
        fm.submit();
    } else {
        alert("모든 항목을 입력하세요.");
    }
}
// 설정 이메일 삭제
function setDel(e) {
    var fm = document.getElementById("setfs");
    //var email = document.getElementById("mail").value;
    var email = $("#setMailList option:selected").val();

    // $("select[name='setMailList'] option:selected").remove();
    // fm.method = "post";
    // fm.target = "_self";
    // fm.action = "/loginTest";
    // fm.submit();

    post("/mailserverDelete", {
        mail: email,
    })

}
// 설정 이메일 수정
function setMod(e) {
    var iAdd = document.getElementById("imap_add").value;
    var iPort = document.getElementById("imap_port").value;
    var sAdd = document.getElementById("smtp_add").value;
    var sPort = document.getElementById("smtp_port").value;
    var mail = document.getElementById("mail").value;
    var mailPw = document.getElementById("mail_passwd").value;
    var tempmail = document.getElementById("tempmail");


    if (iAdd != "" || iPort != "" || sAdd != "" || sPort != "" || mail != "" || mailPw != "" || tempmail != "") {
        var fm = document.getElementById("setf");
        fm.method = "post";
        fm.target = "_self";
        fm.action = "/mailserverUpdate";
        fm.submit();
    } else {
        alert("모든 항목을 입력하세요.");
    }
}

function setDefault(e) {
    //    var mail = document.getElementById("mail").value;
    var mail = document.getElementById("setMailList");
    //alert("|" + mail + "|");
    if (mail.value != "") {
        mail = mail.options[mail.selectedIndex].value;
        var opt = document.createElement("input");
        opt.value = mail;
        opt.type = "hidden";
        opt.name = "opt"
        var fm = document.getElementById("setf");
        fm.appendChild(opt);
        fm.method = "post";
        fm.target = "_self";
        fm.action = "/defaultmailChange";
        fm.submit();
    } else {
        alert("이메일을 선택하세요.");
    }
}

function logoutClick(e) {
    alert("로그인 페이지로 이동합니다.");
}


//메일 쓰기 - 암호화 체크
function write_enck() {
    var enck = document.getElementById("enck");
    var en_sub = document.getElementById("en_sub");
    if (enck.checked == true) {
        document.getElementById("en_sub").style.display = "block";
    } else
        document.getElementById("en_sub").style.display = "none";
}



function pagination() {
    var req_num_row = 15; // 한페이지에 15행
    var $tr = jQuery('tbody tr');
    var total_num_row = $tr.length; //전체 행 개수
    var num_pages = 0;
    if (total_num_row % req_num_row == 0) {
        num_pages = total_num_row / req_num_row;
    }
    if (total_num_row % req_num_row >= 1) {
        num_pages = total_num_row / req_num_row;
        num_pages++;
        num_pages = Math.floor(num_pages++);
    }


    jQuery('.pagination').append("<li><a class=\"prev\">Previous</a></li>");

    for (var i = 1; i <= num_pages; i++) {
        jQuery('.pagination').append("<li><a>" + i + "</a></li>");
        jQuery('.pagination li:nth-child(2)').addClass("active");
        jQuery('.pagination a').addClass("pagination-link");
    }

    jQuery('.pagination').append("<li><a class=\"next\">Next</a></li>");

    $tr.each(function(i) {
        jQuery(this).hide();
        if (i + 1 <= req_num_row) {
            $tr.eq(i).show();
        }
    });

    jQuery('.pagination a').click('.pagination-link', function(e) {
        e.preventDefault();
        $tr.hide();
        var page = jQuery(this).text();
        var temp = page - 1;
        var start = temp * req_num_row;
        var current_link = temp;

        jQuery('.pagination li').removeClass("active");
        jQuery(this).parent().addClass("active");

        for (var i = 0; i < req_num_row; i++) {
            $tr.eq(start + i).show();
        }

        if (temp >= 1) {
            jQuery('.pagination li:first-child').removeClass("disabled");
        } else {
            jQuery('.pagination li:first-child').addClass("disabled");
        }

    });

    jQuery('.prev').click(function(e) {
        e.preventDefault();
        jQuery('.pagination li:first-child').removeClass("active");
    });

    jQuery('.next').click(function(e) {
        e.preventDefault();
        jQuery('.pagination li:last-child').removeClass("active");
    });

}

// jQuery('document').ready(function() {
//     pagination();

//     jQuery('.pagination li:first-child').addClass("disabled");

// });



// $('#pagination').twbsPagination({
//     totalPages: 35,
//     visiblePages: 7,
//     onPageClick: function (event, page) {
//         $('#page-content').text('Page ' + page);
//     }
// });