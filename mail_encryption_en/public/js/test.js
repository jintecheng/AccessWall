window.onload = function () {
  //메뉴 - 사용자 이메일 리스트
  const mailList = document.getElementById("mailList");
  if (mailList) {
    mailListUp();
  }
  // 500 width - hide sidebar
  var alterClass = function () {
    var ww = document.body.clientWidth;
    if (ww < 500) {
      $('.test').removeClass('blue');
      $("body").addClass("sidebar-toggled");
      $(".sidebar").addClass("toggled");

    }
    else if (ww >= 510) {
      $('.test').addClass('blue');
      $("body").removeClass("sidebar-toggled");
      $(".sidebar").removeClass("toggled");
    };
  };

  var cachedWidth = $(window).width();
  $(window).resize(function () {
    var newWidth = $(window).width();
    if (newWidth !== cachedWidth) {
      //PUT YOUR RESIZE HERE
      alterClass();
      cachedWidth = newWidth;
    }
  });
  //Fire it when the page first loads:
  alterClass();

  // 로그인
  const bnt_login = document.getElementById("bnt_login");
  if (bnt_login) bnt_login.addEventListener("click", loginPost, false);

  // 로그아웃
  const bnt_logout = document.getElementById("bnt_logout");
  if (bnt_logout) {
    bnt_logout.addEventListener("click", logoutClick, false);
  }

  // 아이디 찾기
  const forgotIdBtn = document.getElementById("forgotId");
  if (forgotIdBtn) forgotIdBtn.addEventListener("click", forgotIdFunc, false);

  const btnFind = document.getElementById("btnFind");
  if (btnFind) btnFind.addEventListener("click", btnFindFunc, false);

  // 비밀번호 찾기
  const forgotPwBtn = document.getElementById("forgotPw");
  if (forgotPwBtn) forgotPwBtn.addEventListener("click", forgotPwFunc, false);

  const btnPwFind = document.getElementById("btnPwFind");
  if (btnPwFind) btnPwFind.addEventListener("click", btnFindPwFunc, false);

  // 찾기 취소 버튼
  const findCancel = document.querySelectorAll("#findCancel");
  if (findCancel) {
    for (let index = 0; index < findCancel.length; index++) {
      findCancel[index].addEventListener("click", findCancelFunc, false);
    }
  }

  // 가입
  const bntJoin = document.getElementById("bntJoin");
  if (bntJoin) {
    bntJoin.addEventListener("click", register, false);
  }

  // 엔터 활성화
  const jInput = document.querySelectorAll(".jInput");
  if (jInput) {
    for (let index = 0; index < jInput.length; index++) {
      jInput[index].addEventListener("keyup", registerEnter);
    }
  }

  // 설정 - 사용자 이메일 리스트
  const setMail = document.getElementById("setMail");
  if (setMail) setMail.addEventListener("DOMContentLoaded", setmailListUp());

  // 설정 - 이메일 서버 리스트
  const settingServer = document.getElementById("setMailServer");
  if (settingServer)
    settingServer.addEventListener("DOMContentLoaded", setMailServer());

  //설정 - 추가 버튼
  const setAddBtn = document.getElementById("setAdd");
  if (setAddBtn) {
    setAddBtn.addEventListener("click", setAdd, false);
  }

  //설정 - 삭제 버튼
  const setDelBtn = document.getElementById("setDel");
  if (setDelBtn) {
    setDelBtn.addEventListener("click", setDel, false);
  }

  //설정 - 수정 버튼
  const setSaveBtn = document.getElementById("setSave");
  if (setSaveBtn) {
    setSaveBtn.addEventListener("click", setMod, false);
  }

  //설정 - 기본메일설정 버튼
  const setDefaultBtn = document.getElementById("setDefault");
  if (setDefaultBtn) {
    setDefaultBtn.addEventListener("click", setDefault, false);
  }

  //설정 - 사용자 계정 탈퇴 모달
  const deleteAccount = document.getElementById("deleteAccount");
  if (deleteAccount) {
    deleteAccount.addEventListener("click", setModalOpen, false);
  }

  //설정 - 사용자 비밀번호 변경 모달
  const updatePw = document.getElementById("updatePw");
  if (updatePw) {
    updatePw.addEventListener("click", setModalOpen, false);
  }

  //설정 - 비밀번호 변경
  const modalUpdateBtn = document.getElementById("modalUpdateBtn");
  if (modalUpdateBtn) {
    modalUpdateBtn.addEventListener("click", modalUpdateFunc, false);
  }

  //설정 - 사용자 계정 탈퇴
  const modalDelBtn = document.getElementById("modalDelBtn");
  if (modalDelBtn) {
    modalDelBtn.addEventListener("click", modalDeleteFunc, false);
  }

  // 받은메일 삭제 버튼
  const mailDeleteBtn = document.getElementById("mailDelete");
  if (mailDeleteBtn) {
    mailDeleteBtn.addEventListener("click", mailDelete, false);
  }

  // 메일읽기 삭제 버튼
  const mailReadDeleteBtn = document.getElementById("mailReadDelete");
  if (mailReadDeleteBtn) {
    mailReadDeleteBtn.addEventListener("click", mailReadDelete, false);
  }

  // 메일 refresh 버튼
  const mailRefreshBtn = document.getElementById("mailRefresh");
  if (mailRefreshBtn) {
    mailRefreshBtn.addEventListener(
      "click",
      function () {
        window.location.reload();
      },
      false
    );
  }

  // 메일 쓰기, 읽기 암호화 숨기기 버튼
  const mailEnckBtn = document.getElementById("enck");
  if (mailEnckBtn) {
    mailEnckBtn.addEventListener("click", writeEnck, false);
  }

  // 메일 전송 버튼
  const mailWriteBtn = document.getElementById("mailWrite");
  if (mailWriteBtn) {
    mailWriteBtn.addEventListener("click", write, false);
  }

  // pop3 test
  const pop3Btn = document.getElementById("pop3Btn");
  if (pop3Btn) {
    pop3Btn.addEventListener("click", pop3Func, false);
  }
};

function findCancelFunc() {
  let loginDiv = document.querySelector(".loginDiv");
  loginDiv.style.display = "block";

  let findIdDiv = document.querySelector(".finIdDiv");
  findIdDiv.setAttribute('hidden', "true");
  document.getElementById("userName").value = "";
  document.querySelector("#userAddress").value = "";

  let findPwDiv = document.querySelector(".finPwDiv");
  findPwDiv.setAttribute('hidden', "true");
  document.querySelector("#userName2").value = "";
  document.querySelector("#userAddress2").value = "";

  var loginAlert = document.querySelectorAll("#loginAlert")
  loginAlert[0].style.display = "none";
  loginAlert[1].style.display = "none";
  loginAlert[2].style.display = "none";
}

// 아이디 찾기
function forgotIdFunc() {

  let loginDiv = document.querySelector(".loginDiv");
  loginDiv.style.display = "none";

  let findIdDiv = document.querySelector(".finIdDiv");
  findIdDiv.removeAttribute('hidden');
  document.getElementById("id").value = "";
  document.getElementById("pw").value = "";
};

function btnFindFunc() {

  const userName = document.getElementById("userName").value;
  const userAddress = document.getElementById("userAddress").value;
  // let findIdDiv = document.querySelector(".finIdDiv");

  let xhr = new XMLHttpRequest();
  let u3 = {
    Name: userName,
    Email: userAddress,
  };

  xhr.onload = function () {
    if (xhr.status === 200 || xhr.status === 201) {
      alert("Check your email: " + userAddress);
      findCancelFunc();
    } else {
      var loginAlert = document.querySelectorAll("#loginAlert")
      loginAlert[1].style.display = "block";
      $("#userName").focus();
    }
  };

  xhr.open("POST", "/idFind");
  xhr.setRequestHeader("Content-Type", "application/json"); // 컨텐츠타입을 json으로
  //
  xhr.send(JSON.stringify(u3)); // 객체를 JSON 형식의 문자열로 변형
};

// 비밀번호 찾기
function forgotPwFunc() {

  let loginDiv = document.querySelector(".loginDiv");
  loginDiv.style.display = "none";

  let findPwDiv = document.querySelector(".finPwDiv");
  findPwDiv.removeAttribute('hidden');

  document.getElementById("id").value = "";
  document.getElementById("pw").value = "";
};

function btnFindPwFunc() {
  const userAddress = document.getElementById("userAddress2").value;
  var queryString = $("#f3").serialize();

  $.ajax({
    type: 'post',
    url: '/pwFind',
    data: queryString,
    dataType: 'html',
    error: function (xhr, status, error) {
      var loginAlert = document.querySelectorAll("#loginAlert")
      loginAlert[2].style.display = "block";
      $("#userName2").focus();
    },
    success: function (xhr, status) {
      alert("Check your email: " + userAddress);
      findCancelFunc();
    }
  });
};


function setModalOpen() {
  let value = this.value;
  $(document).off('focusin.modal');
  var queryString = $("#setAccoutFm").serialize();

  $.ajax({
    type: 'post',
    url: '/openModalPwCheck',
    data: queryString,
    dataType: 'html',
    error: function (xhr, status, error) {
      if (xhr.status === 400) {
        alert(xhr.status + ": " + xhr.responseText)
      } else {
        alert(xhr.status + "Please retry: " + xhr.responseText)
      }
      document.getElementById("accountPassword").value = "";
    },
    success: function (xhr, status) {
      // 해당 버튼 보이기
      if (value === "Delete Account") {
        document.getElementById("modalUpdateBtn").style.display = "none";
        document.getElementById("setAccPw2").style.display = "none";
        document.getElementById("setAccPw").style.display = "inline";
        document.getElementById("modalDelBtn").style.display = "inline";
      } else if (value === "Change Password") {
        document.getElementById("setAccPw").style.display = "block";
        document.getElementById("setAccPw2").style.display = "block";
        document.getElementById("modalUpdateBtn").style.display = "inline-block";
        document.getElementById("modalDelBtn").style.display = "none";
      }

      //모달 열기
      let setModal = document.getElementById("setModal");
      let span = document.getElementById("mdClose");
      setModal.style.display = "block";

      // When the user clicks on <span> (x), close the modal
      if (span) {
        span.addEventListener("click", function () {
          setModal.style.display = "none";
          document.getElementById("setAccPw").value = "";
          document.getElementById("setAccPw2").value = "";
        }, false);
      }

      // 모달 영역 밖 선택 시 창 닫기
      window.addEventListener("click", function () {
        if (event.target == setModal) {
          setModal.style.display = "none";
          document.getElementById("setAccPw").value = "";
          document.getElementById("setAccPw2").value = "";
        }
      }, false);
    }
  });
}

// 비밀번호 변경
function modalUpdateFunc() {

  var queryString = $("#setModalForm").serialize();

  $.ajax({
    type: 'post',
    url: '/changeAccountPassword',
    data: queryString,
    dataType: 'html',
    error: function (xhr, status, error) {
      if (xhr.status === 400) {
        alert(xhr.status + ": " + xhr.responseText)
        document.getElementById("setAccPw").value = "";
        document.getElementById("setAccPw2").value = "";
      } else {
        alert(xhr.status + " Please retry: " + xhr.responseText)
        document.getElementById("setAccPw").value = "";
        document.getElementById("setAccPw2").value = "";
      }
    },
    success: function (xhr, status) {
      alert("Success Update Password.");
      //모달 닫기
      let setModal = document.getElementById("setModal");
      setModal.style.display = "none";
      document.getElementById("accountPassword").value = "";
      document.getElementById("setAccPw").value = "";
      document.getElementById("setAccPw2").value = "";
    }
  });
};

//계정 삭제
function modalDeleteFunc() {
  var queryString = $("#setModalForm").serialize();

  $.ajax({
    type: 'post',
    url: '/deleteAccount',
    data: queryString,
    dataType: 'html',
    error: function (xhr, status, error) {
      if (xhr.status === 400) {
        alert(xhr.status + ": " + xhr.responseText)
        document.getElementById("setAccPw").value = "";
        document.getElementById("setAccPw2").value = "";
      } else {
        alert(xhr.status + " Please retry: " + xhr.responseText)
        document.getElementById("setAccPw").value = "";
        document.getElementById("setAccPw2").value = "";
      }
    },
    success: function (xhr, status) {
      alert("Success Delete Accout.");
      window.location.href = "/logout";
    }
  });
};


function registerEnter(e) {
  if (e.keyCode === 13) {
    e.preventDefault();
    let joinBtn = document.getElementById("bntJoin");
    let loginBtn = document.getElementById("bnt_login");
    let idFindBtn = document.getElementById("btnFind");
    let pwFindBtn = document.getElementById("btnPwFind");

    if (joinBtn) {
      joinBtn.click();
    }
    if (loginBtn) {
      loginBtn.click();
    }
    if (idFindBtn) {
      idFindBtn.click();
    }
    if (pwFindBtn) {
      pwFindBtn.click();
    }
  }
}

function register() {

  var queryString = $("#joinFm").serialize();
  if (inputCheck() != 0) {
    $.ajax({
      type: 'post',
      url: '/register',
      data: queryString,
      dataType: 'html',
      error: function (xhr, status, error) {
        if (xhr.status === 400) {
          alert("Already existed ID.")
          $("#id").focus();
        } else if (xhr.status === 402) {
          alert("Already existed Mail ID.")
          $("#email").focus();
        } else {
          alert("Register Error: " + error)
        }
      },
      success: function (xhr, status) {
        location.href = "/login";
      }
    });
  }
}

// 단순 submit
function post(URL, PARAMS) {
  let temp = document.createElement("form");
  temp.action = URL;
  temp.method = "post";
  temp.style.display = "none";
  for (let x in PARAMS) {
    let opt = document.createElement("textarea");
    opt.name = x;
    opt.value = PARAMS[x];
    temp.appendChild(opt);
  }
  document.body.appendChild(temp);
  temp.submit();
}

//메일 목록에서 보낸 사람 클릭 시 메일쓰기로 이동
function page_move(s_page) {
  post("/compose", {
    bo: s_page,
  });
}

//메일 보내기
function write() {
  const to = document.getElementById("writeTo").value;
  // var cc = document.getElementById("write_cc").value;
  const title = document.getElementById("enTitle").value;
  const input = document.getElementById("enSource").value;

  let xhr = new XMLHttpRequest();
  let data = {
    jto: to,
    //jcc: cc,
    jtitle: title,
    jinput: input,
  };
  xhr.onload = function () {
    if (xhr.status === 200 || xhr.status === 201) {
      //console.log(xhr.responseText);
      alert("Completed Send.");
      window.history.back();
    } else {
      //console.error(xhr.responseText);
      alert("Error");
      window.location.reload();
    }
  };
  xhr.open("POST", "/write");
  xhr.setRequestHeader("Content-Type", "application/json"); // 컨텐츠타입을 json으로
  //
  xhr.send(JSON.stringify(data)); // 객체를 JSON 형식의 문자열로 변형
  // xhr.abort(); // 전송된 요청 취소
};

//로그인
function loginPost() {
  var queryString = $("#f1").serialize();
  $.ajax({
    type: 'post',
    url: '/loginTest',
    data: queryString,
    dataType: 'html',
    error: function (xhr, status, error) {
      var loginAlert = document.querySelectorAll("#loginAlert")
      loginAlert[0].style.display = "block";
      $("#id").focus();
    },
    success: function (xhr, status) {
      //  console.log("성공");
      location.href = "/index";
    }
  });
}

function pop3Func() {
  $.ajax({
    type: 'get',
    url: '/pop',
    //data: queryString,
    dataType: 'html',
    error: function (xhr, status, error) {
      console.log("실패");

    },
    success: function (data, status) {
      //var json = $.parseJSON(data);
      console.log("data: \n" + data);
      const data1 = JSON.parse(data);
      console.log("data1: \n" + data1);
      const title = data1[8].Title;
      const from = data1[8].From;
      const to = data1[8].To;
      const date = data1[8].Date;

      console.log("1Title: " + title);
      console.log("1From: " + from);
      console.log("1To: " + to);
      console.log("1Date: " + date);

      $('#table-responsive').parent().prepend('<label>Date: ' + date + '</label><br>');
      $('#table-responsive').parent().prepend('<label>To: ' + to + '</label><br>');
      $('#table-responsive').parent().prepend('<label>From: ' + from + '</label><br>');
      $('#table-responsive').parent().prepend('<label>Title: ' + title + '</label><br>');

      document.getElementById("table-responsive").innerHTML = data1[13].Content;

      for (var ele in data1) {
        console.log("--------------------------");
        console.log("Title: " + data1[ele].Title);
        console.log("From: " + data1[ele].From);
        console.log("To: " + data1[ele].To);
        console.log("Date: " + data1[ele].Date);
      }
    }
  });
}

function download(filename, text) {
  var element = document.createElement('a');
  element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
  element.setAttribute('download', filename);

  element.style.display = 'none';
  document.body.appendChild(element);

  element.click();

  document.body.removeChild(element);
}

// 이전 페이지 돌아가기
function goBack() {
  window.history.back();
}

function mailDelete() {
  var result = confirm("Are you sure you want to delete it?");
  if (result) {
    const menuNum = document.getElementById("menuNum"); // 메일함 종류
    const mailChk = document.getElementsByName("mailChk");
    let checked = 0;

    let a = new Array();

    for (let i = 0; i < mailChk.length; i++) {
      if (mailChk[i].checked) {
        a.push(mailChk[i].value);
        checked++;
      }
    }
    if (checked === 0) alert("Please select mail.");
    else {
      let xhr = new XMLHttpRequest();

      let data = {
        Mail: a,
        Num: menuNum.value,
      };
      xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201) {
          for (let i = 0; i < a.length; i++) {
            $("#tr" + a[i]).remove();
          }
          alert("Delete Complete.");
        } else {
          console.error(xhr.responseText);
          alert("Error");
          window.location.reload();
        }
      };
      xhr.open("POST", "/mailDelete");
      xhr.setRequestHeader("Content-Type", "application/json"); // 컨텐츠타입을 json으로
      xhr.send(JSON.stringify(data)); // 데이터를 stringify해서 보냄
      // xhr.abort(); // 전송된 요청 취소
    }
  } else {
    alert("Cancel");
  }
}

// 메일 읽기 삭제
function mailReadDelete() {
  var result = confirm("Are you sure you want to delete it?");
  if (result) {
    const id = document.getElementById("mailReadDelete").value; // 메일함 종류
    const num = document.getElementById("readPageNum").value; // 메일함 종류
    //console.log("id: " + id);
    //console.log("num: " + num);

    let a = new Array();
    a.push(id);

    let xhr = new XMLHttpRequest();

    let data = {
      Mail: a,
      Num: num,
    };

    //console.log("data a 확인: "+data.a);
    xhr.onload = function () {
      if (xhr.status === 200 || xhr.status === 201) {
        alert("Delete Complete.");
        window.location.href = "/index"
      } else {
        console.error(xhr.responseText);
        alert("Error");
        window.location.reload();
      }
    };
    xhr.open("POST", "/mailDelete");
    xhr.setRequestHeader("Content-Type", "application/json"); // 컨텐츠타입을 json으로
    xhr.send(JSON.stringify(data)); // 데이터를 stringify해서 보냄

  } else {
    alert("Cancel");
  }
}

// 설정 - 이메일 서버 셀렉트 선택 시 주소와 포트번호 자동 설정 setonchange
function setEmailServer(e) {
  let chk = document.querySelector('input[name="chk_info"]:checked').value;
  //console.log(chk); //imap, pop
  //console.log(e); // 메일 서버

  let ia = document.getElementById("imap_add");
  let ip = document.getElementById("imap_port");
  let sa = document.getElementById("smtp_add");
  let sp = document.getElementById("smtp_port");
  // var ma = document.getElementById("mail");
  // var pw = document.getElementById("mail_passwd");

  let xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status === 200 || xhr.status === 201) {
      search1 = xhr.response;
      //console.log("1" + search1);
      var obj = JSON.parse(search1);
      //console.log("2" + obj.imap_add);
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
      alert("Error");
      window.location.reload();
    }
  };
  xhr.open("GET", "/modMailServer?chk=" + chk + "&mail=" + e, true);
  xhr.send();
}

//메뉴 - 사용자 이메일 리스트 가져오기
function mailListUp() {
  let xhr = new XMLHttpRequest();
  //alert(sessionStorage.getItem('eid'));
  //console.log("in js pw: " + pw);
  xhr.onload = function () {
    if (xhr.status === 200 || xhr.status === 201) {
      //console.log(xhr.responseText);
      search = xhr.responseText;
      let jbSplit = search.split(";");
      jbSplit.length = jbSplit.length - 1;
      //  $('#mailList').empty();
      //console.log("이메일 리스트: " + search);
      for (let i in jbSplit) {
        let addOp = document.createElement("option");
        // addOp.setAttribute("value", i + 1);
        addOp.innerHTML = jbSplit[i] + "</option>";
        document.getElementById("mailList").appendChild(addOp);
        //console.log(i + ": " + jbSplit[i]);
        if (i + 1 == jbSplit.length) break;
      }
      //console.log("mailListVal in Up: " + sessionStorage.getItem('eid'));
      if (sessionStorage.getItem("eid") != null) {
        for (
          i = 0;
          i < document.getElementById("mailList").options.length;
          i++
        ) {
          //console.log(sessionStorage.getItem("eid"));
          if (
            document.getElementById("mailList").options[i].value ==
            sessionStorage.getItem("eid")
          ) {
            document.getElementById("mailList").options[i].selected =
              "selected";
          }
        }
      }
      // alert("전송 완료!");
    } else {
      //console.error(xhr.responseText);
      alert("Error");
      window.location.reload();
    }
  };

  xhr.open("POST", "/mailList", true);
  xhr.setRequestHeader("Content-Type", "text/plain");
  xhr.send();
}

//설정 - 사용자 이메일 리스트 가져오기
function setmailListUp() {
  let xhr = new XMLHttpRequest();

  // console.log("in js pw: " + pw);
  xhr.onload = function () {
    if (xhr.status === 200 || xhr.status === 201) {
      search = xhr.responseText;
      let jbSplit = search.split(";");
      jbSplit.length = jbSplit.length - 1;
      //  $('#mailList').empty();
      ///console.log("설정 이메일 서버1" + search);
      for (let i in jbSplit) {
        let addOp = document.createElement("option");
        addOp.setAttribute("value", jbSplit[i]);
        addOp.innerHTML = jbSplit[i] + "</option>";
        document.getElementById("setMailList").appendChild(addOp);
        //console.log(i + ": " + jbSplit[i]);
        if (i + 1 == jbSplit.length) break;
      }
    } else {
      //console.error(xhr.responseText);
      alert("Error");
      window.location.reload();
    }
  };
  xhr.open("POST", "/mailList", true);
  xhr.setRequestHeader("Content-Type", "text/plain");
  xhr.send();
}

//설정 - 이메일 서버 리스트 가져오기
function setMailServer() {
  let xhr = new XMLHttpRequest();
  //    console.log("in js pw: " + pw);
  xhr.onload = function () {
    if (xhr.status === 200 || xhr.status === 201) {
      search = xhr.responseText;
      //console.log("설정 이메일 서버2" + search);
      let jbSplit = search.split(";");
      jbSplit.length = jbSplit.length - 1;

      for (var i in jbSplit) {
        var addOp = document.createElement("option");
        addOp.setAttribute("value", jbSplit[i]);
        addOp.innerHTML = jbSplit[i] + "</option>";
        document.getElementById("setMailServer").appendChild(addOp);
        //console.log("설정 이메일 서버12" + i + ": " + jbSplit[i]);
        if (i + 1 == jbSplit.length) break;
      }
      // alert("전송 완료!");
    } else {
      //console.error(xhr.responseText);
      alert("Error");
      window.location.reload();
    }
  };
  xhr.open("POST", "/mailPreset", true);
  xhr.setRequestHeader("Content-Type", "text/plain");
  xhr.send();
}

//메뉴 - 사용자 메일리스트 선택
function mailList_ch(e) {
  let xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function () {
    if (this.readyState === 4 && this.status === 200) {
      //console.log(xhr.responseText);
      search = xhr.responseText;

      //xhr.response
      let jbSplit = search.split(";");
      jbSplit.length = jbSplit.length - 1;
      //  $('#mailList').empty();

      for (let i in jbSplit) {
        let addOp = document.createElement("option");
        addOp.innerHTML = jbSplit[i] + "</option>";
        document.getElementById("mailList").appendChild(addOp);
        //console.log(i + ": " + jbSplit[i]);
        if (i + 1 == jbSplit.length) break;
      }
      sessionStorage.setItem("eid", e);
      //console.log("eid in Ch: " + sessionStorage.getItem('eid'));
      if (search == "1") {
        location.replace("/index");
      } else {
        location.reload();
      }
    }
  };
  let para = document.location.href.split("/");
  //console.log(para);

  xhr.open("GET", "/mailChange?id=" + e + "&url=" + para[3], true);
  xhr.send();
}

//  설정 페이지 - 메일 리스트 클릭
function setListClick(e) {
  //console.log("this.value e: " + e)
  let ia = document.getElementById("imap_add");
  let ip = document.getElementById("imap_port");
  let sa = document.getElementById("smtp_add");
  let sp = document.getElementById("smtp_port");
  let ma = document.getElementById("mail");
  let pw = document.getElementById("mail_passwd");
  let tempmail = document.getElementById("tempmail");

  var xhr = new XMLHttpRequest();

  xhr.onreadystatechange = function () {
    if (this.readyState === 4 && this.status === 200) {
      //console.log(xhr.responseText);

      search1 = xhr.response;
      //console.log(search1);
      let obj = JSON.parse(search1);
      // console.log(obj.imap_add);

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

//  설정 페이지 - 이메일 추가
function setAdd() {
  const iAdd = document.getElementById("imap_add").value;
  const iPort = document.getElementById("imap_port").value;
  const sAdd = document.getElementById("smtp_add").value;
  const sPort = document.getElementById("smtp_port").value;
  const mail = document.getElementById("mail").value;
  const mailPw = document.getElementById("mail_passwd").value;
  let options = $('#setMailList').find('option').map(function () { return $(this).val(); }).get();

  if (
    iAdd !== "" &&
    iPort !== "" &&
    sAdd !== "" &&
    sPort !== "" &&
    mail !== "" &&
    mailPw !== ""
  ) {
    if (options.includes(mail) == true) {
      alert("The same email address already exists.");
    } else {
      let fm = document.getElementById("setf");
      fm.method = "post";
      fm.target = "_self";
      fm.action = "/mailserverInsert";
      fm.submit();
    }
  } else {
    alert("Please enter all items.");
  }
  let plusUl = document.createElement("option");
  plusUl.innerHTML = "새로운 이메일</option>";
  document.getElementById("mailList").appendChild(plusUl);
}

//  설정 페이지 - 이메일 삭제
function setDel() {
  //var fm = document.getElementById("setfs");
  //var email = document.getElementById("mail").value;
  const email = $("#setMailList option:selected").val();
  let x = document.getElementById("setMailList").selectedIndex;
  let y = document.getElementById("setMailList").options;
  let idx = y[x].index;
  let val = y[x].value;

  if (idx == 0) {
    alert("Default mail cannot be deleted.");
  } else {

    // $("select[name='setMailList'] option:selected").remove();
    // fm.method = "post";
    // fm.target = "_self";
    // fm.action = "/loginTest";
    // fm.submit();
    post("/mailserverDelete", {
      mail: val,
    });
  }
}

//  설정 페이지 - 이메일 수정
function setMod() {
  const iAdd = document.getElementById("imap_add").value;
  const iPort = document.getElementById("imap_port").value;
  const sAdd = document.getElementById("smtp_add").value;
  const sPort = document.getElementById("smtp_port").value;
  const mail = document.getElementById("mail").value;
  const mailPw = document.getElementById("mail_passwd").value;
  const tempmail = document.getElementById("tempmail");

  if (
    iAdd !== "" &&
    iPort !== "" &&
    sAdd !== "" &&
    sPort !== "" &&
    mail !== "" &&
    mailPw !== "" &&
    tempmail !== ""
  ) {
    let fm = document.getElementById("setf");
    fm.method = "post";
    fm.target = "_self";
    fm.action = "/mailserverUpdate";
    fm.submit();
  } else {
    alert("Please enter all items.");
  }
}

// 설정 페이지 - 메일 기본값 설정
function setDefault() {
  //    var mail = document.getElementById("mail").value;
  let mail = document.getElementById("setMailList");
  //alert("|" + mail + "|");
  if (mail.value != "") {
    mail = mail.options[mail.selectedIndex].value;
    let opt = document.createElement("input");
    opt.value = mail;
    opt.type = "hidden";
    opt.name = "opt";
    let fm = document.getElementById("setf");
    fm.appendChild(opt);
    fm.method = "post";
    fm.target = "_self";
    fm.action = "/defaultmailChange";
    fm.submit();
  } else {
    alert("Please select mail.");
  }
}

function logoutClick() {
  sessionStorage.removeItem("eid");
  //alert("Logout.");
  window.location.href = "/logout";
}

//메일 쓰기 - 암호화 체크
// const writeEnck = () => {
function writeEnck() {
  const enck = document.getElementById("enck");
  //var en_sub = document.getElementById("en_sub");
  if (enck.checked) {
    document.getElementById("en_sub").style.display = "block";
  } else document.getElementById("en_sub").style.display = "none";
};


// 회원가입 빈칸
// $(document).ready(function () {
//   $("#bntJoin").click(function () {
function inputCheck() {
  if ($("#name").val().length == 0) {
    alert("Please input your name.");
    $("#name").focus();
    return 0;
  }
  if ($("#id").val().length == 0) {
    alert("Please input your ID.");
    $("#id").focus();
    return 0;
  }
  if ($("#password1").val().length == 0) {
    alert("Please input your Password.");
    $("#password1").focus();
    return 0;
  }
  if ($("#password2").val().length == 0) {
    alert("Please input your Password.");
    $("#password2").focus();
    return 0;
  }
  if ($("#smtp_add").val().length == 0) {
    alert("Please input your SMTP Address.");
    $("#smtp_add").focus();
    return 0;
  }
  if ($("#smtp_port").val().length == 0) {
    alert("Please input your SMTP Port.");
    $("#smtp_port").focus();
    return 0;
  }
  if ($("#imap_add").val().length == 0) {
    alert("Please input your IMAP Address.");
    $("#imap_add").focus();
    return 0;
  }
  if ($("#imap_port").val().length == 0) {
    alert("Please input your IMAP Port.");
    $("#imap_port").focus();
    return 0;
  }
  if ($("#email").val().length == 0) {
    alert("Please input your Email Address.");
    $("#email").focus();
    return 0;
  }
  if ($("#mail_passwd").val().length == 0) {
    alert("Please input your Email Password.");
    $("#mail_passwd").focus();
    return 0;
  }
  if ($("#password2").val() != $("#password1").val()) {
    alert("Please check your password.");
    $("#password2").focus();
    return 0;
  }
  return 1;
}