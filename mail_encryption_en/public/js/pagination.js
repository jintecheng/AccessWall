function pagination(a) {
    var req_num_row = 10; //한 페이지에 출력할 행 개수
    var $tr = jQuery('.tb tr');
    var total_num_row = $tr.length - a; //테이블 총 행
    var num_pages = 0; //페이지 번호 개수
    var current_page = 0;
    if ($(window).width() < 500) {
        var visiblePages = 4;
    } else {
        var visiblePages = 5;
    }
    var visibleTemp = visiblePages;

    // 테이블 전체행과 한페이지에 출력할 행 개수 비교
    if (total_num_row % req_num_row == 0) {
        num_pages = total_num_row / req_num_row;
        var endPage = num_pages;
    }
    if (total_num_row % req_num_row >= 1) {
        num_pages = total_num_row / req_num_row;
        num_pages++;
        num_pages = Math.floor(num_pages++);
        var endPage = num_pages;
    }
    if (num_pages <= visiblePages) var num = num_pages; else var num = visiblePages;

    // 아래 두 줄은 동일
    // $('div').css('border', '4px solid #f00');
    // jQuery('div').css('border', '4px solid #f00');
    // $('div') 의미는 'div'를 매개변수 값으로 $() 함수를 호출한 것입니다.


    //pagination 클래스 아래에 번호와 버튼들 추가
    jQuery('.pagination').append("<li class=\"page-item disabled\"><a class=\"page-link first\">First</a></li>");
    jQuery('.pagination').append("<li class=\"page-item\"><a class=\"page-link prev\" aria-label=\"Previous\"><span aria-hidden=\"true\">&laquo;</span></a></li>");

    for (var i = 1; i <= num; i++) {
        jQuery('.pagination').append("<li class=\"page-item\"><a class=\"page-link page-num\">" + i + "</a></li>");
        jQuery('.pagination li:nth-child(3)').addClass("active");
        jQuery('.page-num').addClass("pagination-link");
    }

    jQuery('.pagination').append("<li class=\"page-item\"><a class=\"page-link next\" aria-label=\"Next\"><span aria-hidden=\"true\">&raquo;</span></a></li>");
    jQuery('.pagination').append("<li class=\"page-item\"><a class=\"page-link last\">Last</a></li>");

    //tr: 테이블의 한줄 전체, 모든 테이블 행을 숨기고 첫번째 페이지에 뜰 화면만 출력
    $tr.each(function (i) {
        jQuery(this).hide();
        if (i + 1 <= req_num_row) {
            $tr.eq(i).show();
        }
    });

    //페이지가 하나이거나 없으면 모든 버튼 비활성화
    if (num_pages == 1 || num_pages == 0) {
        jQuery('.pagination li:first-child').addClass("disabled");
        jQuery('.pagination li:first-child').next().addClass("disabled");
        jQuery('.pagination li:last-child').addClass("disabled");
        jQuery('.pagination li:last-child').prev().addClass("disabled");
        jQuery('.pagination li:first-child').next().next().addClass("disabled");
    }

    //페이지 번호 클릭
    jQuery('.page-num').click('.pagination-link', function (e) {
        e.preventDefault();
        $tr.hide(); //요소 사라짐
        var page = parseInt(jQuery(this).text()); //<a class="page-link pagination-link">2</a> -> 텍스트 2
        //console.log("page: "+page);
        var temp = page - 1;
        var start = temp * req_num_row;
        current_page = temp;

        // console.log("currentPage: "+current_page+" / visibleTemp: "+visibleTemp)
        jQuery('.pagination li').removeClass("active");
        jQuery(this).parent().addClass("active"); //a말고 li에 active

        for (var i = 0; i < req_num_row; i++) {
            $tr.eq(start + i).show();
        }
        chkDisabled(temp);
    });

    //first 버튼 클릭
    jQuery('.first').click(function (e) {
        e.preventDefault();
        $tr.hide(); //요소 사라짐
        var temp = 0;
        current_page = temp;
        visibleTemp = visiblePages;
        var start = temp * req_num_row;
        jQuery('.pagination li').removeClass("active");
        // jQuery(this).parent().addClass("active");

        for (var i = 0; i < req_num_row; i++) {
            $tr.eq(start + i).show();
        }
        for (let index = 0; index < visiblePages; index++) {
            // if(current_page+index>=endPage){
            jQuery('.page-num').eq(index).show();
            // }else{
            //  console.log(current_page);
            jQuery('.page-num').eq(index).text(index + 1);
            //  }
        }
        jQuery('.pagination li:first-child').next().next().addClass("active");
        jQuery('.pagination li:first-child').removeClass("active");

        chkDisabled(temp);
    });
    //last 버튼 클릭
    jQuery('.last').click(function (e) {
        $tr.hide(); //요소 사라짐
        var temp = num_pages - 1;
        current_page = temp;
        var start = current_page * req_num_row;

        //console.log("start: "+endPage +" / visiblePages: "+visiblePages+" / current: "+current_page+" / temp: " + temp);

        jQuery('.pagination li').removeClass("active");
        // jQuery(this).parent().addClass("active");

        for (var i = 0; i < req_num_row; i++) {
            $tr.eq(start + i).show();
        }

        var cm1 = (temp + 1) % visiblePages; // 마지막페이지 나누기 범위의 나머지
        var cm2 = Math.floor(temp / visiblePages); // 마지막페이지 나누기 범위의 몫

        // console.log("cm1: "+cm1+" / cm2: "+cm2);

        if (cm1 == 0) {
            var i = 0;
            for (let index = visiblePages - 1; index >= 0; index--) {
                jQuery('.page-num').eq(index).text(num_pages - i);
                jQuery('.pagination li').eq(visiblePages + 1).addClass("active");
                i++;
            }
        } else {
            for (let index = 0; index < visiblePages; index++) {

                if (index >= visiblePages - (visiblePages - cm1)) {
                    //console.log("index: "+index+" |visi: "+ visiblePages+" |cm1: "+cm1);
                    jQuery('.page-num').eq(index).hide();
                } else {
                    // console.log("c_p: "+current_page+" / endpage: "+endPage);
                    //                    jQuery('.page-num').eq(index).text(current_page+index+1);
                    jQuery('.page-num').eq(index).text(index + 1 + (visiblePages * cm2));
                    if (current_page + index + 1 == endPage) jQuery('.pagination li').eq(cm1 + 1).addClass("active");
                }
            }
        }

        visibleTemp = Math.floor(visiblePages * (cm2 + 1));
        //console.log("vT: "+visibleTemp);
        // jQuery('.pagination li:last-child').prev().prev().addClass("active");
        jQuery('.pagination li:first-child').removeClass("active");

        chkDisabled(temp);
    });
    // 이전 버튼 클릭
    jQuery('.prev').click(function (e) {
        e.preventDefault();
        $tr.hide(); //요소 사라짐
        var temp = current_page - 1;
        var start = temp * req_num_row;
        jQuery('.pagination li').removeClass("active");
        jQuery('.page-num').eq(temp % visiblePages).parent().addClass("active");
        //console.log("currentPage: "+current_page+" / visibleTemp: "+visibleTemp)

        for (var i = 0; i < req_num_row; i++) {
            $tr.eq(start + i).show();
        }
        if (current_page == visibleTemp - visiblePages && visibleTemp != visiblePages) {
            //console.log("123currentPage: "+current_page+" / visibleTemp: "+visibleTemp)
            visibleTemp = visibleTemp - visiblePages;
            var i = 0;
            for (let index = visiblePages - 1; index >= 0; index--) {
                jQuery('.page-num').eq(index).show();
                jQuery('.page-num').eq(index).text(current_page - i);
                i++;
            }
        }
        current_page = temp;
        chkDisabled(temp);
    });

    // 넥스트했을때 visible을 넘어가면 5개의 번호 갱신, visible 개수보다 적으면 삭제(반복문 prev prev)으로 
    jQuery('.next').click(function (e) {
        e.preventDefault();
        $tr.hide(); //요소 사라짐
        var temp = current_page + 1;
        current_page = temp;
        var start = temp * req_num_row;
        //console.log("currentPage: "+current_page+" / visibleTemp: "+visibleTemp)
        jQuery('.pagination li').removeClass("active");
        jQuery('.page-num').eq(current_page % visiblePages).parent().addClass("active");

        for (var i = 0; i < req_num_row; i++) {
            $tr.eq(start + i).show();
        }
        if (current_page == visibleTemp) {
            visibleTemp = visibleTemp + visiblePages;
            for (let index = 0; index < visiblePages; index++) {
                if (current_page + index >= endPage) {
                    jQuery('.page-num').eq(index).hide();
                } else {
                    //  console.log(current_page);
                    jQuery('.page-num').eq(index).text(index + current_page + 1);
                }
            }
        }
        chkDisabled(temp);
    });

    // 현재 페이지에 따라 버튼들의 비활성화 여부 적용
    function chkDisabled(current_num) {
        if (current_num == 0) {
            jQuery('.pagination li:first-child').addClass("disabled");
            jQuery('.pagination li:first-child').next().addClass("disabled");
            jQuery('.pagination li:last-child').removeClass("disabled");
            jQuery('.pagination li:last-child').prev().removeClass("disabled");
            //console.log(current_num+ " 1/ " + num_pages)
        } else if (current_num == num_pages - 1) {
            // console.log(current_num + " 2/ "+num_pages)
            jQuery('.pagination li:last-child').addClass("disabled");
            jQuery('.pagination li:last-child').prev().addClass("disabled");
            jQuery('.pagination li:first-child').removeClass("disabled");
            jQuery('.pagination li:first-child').next().removeClass("disabled");
        } else {
            //console.log(current_num+ " 3/ " + num_pages)
            jQuery('.pagination li:last-child').removeClass("disabled");
            jQuery('.pagination li:last-child').prev().removeClass("disabled");
            jQuery('.pagination li:first-child').removeClass("disabled");
            jQuery('.pagination li:first-child').next().removeClass("disabled");
        }
    }
}

jQuery('document').ready(function () {
    pagination(0);
    jQuery('.pagination li:first-child').addClass("disabled");
    jQuery('.pagination li:first-child').next().addClass("disabled");
});