function AjaxNewsCommonListGet(n) {
    if (noLoadComment) return ! 1;
    noLoadComment = !0;
    $.ajax({
        type: "get",
        url: Api.Comment.NewsCommonListGet,
        data: {
            NewsID: newsID,
            LapinID: lapinID,
            MaxCommentID: maxCommentID,
            Latest: latest
        },
        dataType: "json",
        success: function(t) {
            var f, e, i, u, r;
            if (t.Success === 1) {
                if (t = t.Result, n) {
                    if (firstLoadComment = !n, $.type(t.Tlist) !== "undefined" && t.Tlist !== null && t.Tlist !== "" && t.Tlist.length > 0) {
                        $(".comment .top-comment-box").removeClass("hide");
                        for (i in t.Tlist) {
                            f = '<li class="placeholder main-floor" data-comment-id="' + t.Tlist[i].M.Ci + '"><div class="user"><a title="软媒通行证数字ID：' + t.Tlist[i].M.Ui + '" href="' + t.Tlist[i].M.UserIndexUrl + '"><div class="user-hd"><img src="' + t.Tlist[i].M.HeadImg + '" width="45" height="45" onerror=\'this.src = "' + Domain.WapImageDomain + '/user/noavatar.png"\' /><\/div><span class="lv">Lv.' + t.Tlist[i].M.Ul + '<\/span><\/a><\/div><div class="review-con"><div class="user-mes">';
                            switch (t.Tlist[i].M.M) {
                            case 1:
                                f += '<a class="rmpvipblue" title = "软媒通行证数字ID：' + t.Tlist[i].M.Ui + '" target = "_blank" href = "' + t.Tlist[i].M.UserIndexUrl + '" ><span class="user-name">' + t.Tlist[i].M.N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_blue.svg" width="16" height="15" /><\/a>';
                                break;
                            case 2:
                                f += '<a class="rmpvipred" title="软媒通行证数字ID：' + t.Tlist[i].M.Ui + '" href="' + t.Tlist[i].M.UserIndexUrl + '"><span class="user-name">' + t.Tlist[i].M.N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_itred.svg" width="16" height="15" /><\/a>';
                                break;
                            case 9:
                                f += '<a class="rmpvip" title="软媒通行证数字ID：' + t.Tlist[i].M.Ui + '" href="' + t.Tlist[i].M.UserIndexUrl + '"><span class="user-name">' + t.Tlist[i].M.N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_orange.svg" width="16" height="15" /><\/a>';
                                break;
                            default:
                                f += '<a title="软媒通行证数字ID：' + t.Tlist[i].M.Ui + '" href="' + t.Tlist[i].M.UserIndexUrl + '"><span class="user-name">' + t.Tlist[i].M.N + "<\/span><\/a>"
                            }
                            f += '<span class="mobile ' + t.Tlist[i].M.ClName + '"><a target="_blank" href="' + Domain.WapDomain + '/html/app/appdown.html">' + t.Tlist[i].M.Ta + '<\/a><\/span><span class="user-floor">' + t.Tlist[i].M.SF + '<\/span><\/div><div class="user-write-msg">' + (t.Tlist[i].M.Y !== "" && t.Tlist[i].M.Y !== null ? '<span class="user-ip">' + t.Tlist[i].M.Y + "<\/span>": "") + '<span class="user-write-time">' + t.Tlist[i].M.WT + '<\/span><\/div><div class="user-review">' + t.Tlist[i].M.C + '<\/div><div class="review-footer"><span class="collapse" data-collapsed="" role="button">' + (t.Tlist[i].Hfc > 0 ? "展开(" + t.Tlist[i].Hfc + ")": "") + '<\/span><span class="review-ft-fr"><span class="stand-by" role="button">支持(' + t.Tlist[i].M.S + ')<\/span><span class="oppose" role="button">反对(' + t.Tlist[i].M.A + ')<\/span><span class="reply" role="button">回复<\/span><\/span><\/div><\/div><\/li>';
                            $(".comment .top-comment-box .top-comment").append(f)
                        }
                    }
                    if ($.type(t.Hlist) !== "undefined" && t.Hlist !== null && t.Hlist !== "" && t.Hlist.length > 0) {
                        $(".comment .hot-comment-box").removeClass("hide");
                        for (i in t.Hlist) {
                            e = '<li class="placeholder main-floor" data-comment-id="' + t.Hlist[i].M.Ci + '"><div class="user"><a title="软媒通行证数字ID：' + t.Hlist[i].M.Ui + '" href="' + t.Hlist[i].M.UserIndexUrl + '"><div class="user-hd"><img src="' + t.Hlist[i].M.HeadImg + '" width="45" height="45" onerror=\'this.src = "' + Domain.WapImageDomain + '/user/noavatar.png"\' /><\/div><span class="lv">Lv.' + t.Hlist[i].M.Ul + '<\/span><\/a><\/div><div class="review-con"><div class="user-mes">';
                            switch (t.Hlist[i].M.M) {
                            case 1:
                                e += '<a class="rmpvipblue" title = "软媒通行证数字ID：' + t.Hlist[i].M.Ui + '" target = "_blank" href = "' + t.Hlist[i].M.UserIndexUrl + '" ><span class="user-name">' + t.Hlist[i].M.N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_blue.svg" width="16" height="15" /><\/a>';
                                break;
                            case 2:
                                e += '<a class="rmpvipred" title="软媒通行证数字ID：' + t.Hlist[i].M.Ui + '" href="' + t.Hlist[i].M.UserIndexUrl + '"><span class="user-name">' + t.Hlist[i].M.N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_itred.svg" width="16" height="15" /><\/a>';
                                break;
                            case 9:
                                e += '<a class="rmpvip" title="软媒通行证数字ID：' + t.Hlist[i].M.Ui + '" href="' + t.Hlist[i].M.UserIndexUrl + '"><span class="user-name">' + t.Hlist[i].M.N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_orange.svg" width="16" height="15" /><\/a>';
                                break;
                            default:
                                e += '<a title="软媒通行证数字ID：' + t.Hlist[i].M.Ui + '" href="' + t.Hlist[i].M.UserIndexUrl + '"><span class="user-name">' + t.Hlist[i].M.N + "<\/span><\/a>"
                            }
                            e += '<span class="mobile ' + t.Hlist[i].M.ClName + '"><a target="_blank" href="' + Domain.WapDomain + '/html/app/appdown.html">' + t.Hlist[i].M.Ta + '<\/a><\/span><span class="user-floor">' + t.Hlist[i].M.SF + '<\/span><\/div><div class="user-write-msg">' + (t.Hlist[i].M.Y !== "" && t.Hlist[i].M.Y !== null ? '<span class="user-ip">' + t.Hlist[i].M.Y + "<\/span>": "") + '<span class="user-write-time">' + t.Hlist[i].M.WT + '<\/span><\/div><div class="user-review">' + t.Hlist[i].M.C + '<\/div><div class="review-footer"><span class="collapse" data-collapsed="" role="button">' + (t.Hlist[i].Hfc > 0 ? "展开(" + t.Hlist[i].Hfc + ")": "") + '<\/span><span class="review-ft-fr"><span class="stand-by" role="button">支持(' + t.Hlist[i].M.S + ')<\/span><span class="oppose" role="button">反对(' + t.Hlist[i].M.A + ')<\/span><span class="reply" role="button">回复<\/span><\/span><\/div><\/div><\/li>';
                            $(".comment .hot-comment-box .hot-comment").append(e)
                        }
                    }
                }
                if ($.type(t.Clist) !== "undefined" && t.Clist !== null && t.Clist !== "" && t.Clist.length > 0) {
                    $(".comment .all-comment-box .tags-sort").removeClass("hide");
                    for (i in t.Clist) {
                        maxCommentID = t.Clist[i].M.Ci;
                        u = '<li class="placeholder main-floor" data-comment-id="' + t.Clist[i].M.Ci + '"><div class="user"><a title="软媒通行证数字ID：' + t.Clist[i].M.Ui + '" href="' + t.Clist[i].M.UserIndexUrl + '"><div class="user-hd"><img src="' + t.Clist[i].M.HeadImg + '" width="45" height="45" onerror="this.src=\'' + Domain.WapImageDomain + '/user/noavatar.png\'" /><\/div><span class="lv">Lv.' + t.Clist[i].M.Ul + '<\/span><\/a><\/div><div class="review-con"><div class="user-mes">';
                        switch (t.Clist[i].M.M) {
                        case 1:
                            u += '<a class="rmpvipblue" title="软媒通行证数字ID：' + t.Clist[i].M.Ui + '" href="' + t.Clist[i].M.UserIndexUrl + '"><span class="user-name">' + t.Clist[i].M.N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_blue.svg" width="16" height="15" /><\/a>';
                            break;
                        case 2:
                            u += '<a class="rmpvipred" title="软媒通行证数字ID：' + t.Clist[i].M.Ui + '" href="' + t.Clist[i].M.UserIndexUrl + '"><span class="user-name">' + t.Clist[i].M.N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_itred.svg" width="16" height="15" /><\/a>';
                            break;
                        case 9:
                            u += '<a class="rmpvip" title="软媒通行证数字ID：' + t.Clist[i].M.Ui + '" href="' + t.Clist[i].M.UserIndexUrl + '"><span class="user-name">' + t.Clist[i].M.N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_orange.svg" width="16" height="15" /><\/a>';
                            break;
                        default:
                            u += '<a title="软媒通行证数字ID：' + t.Clist[i].M.Ui + '" href="' + t.Clist[i].M.UserIndexUrl + '"><span class="user-name">' + t.Clist[i].M.N + "<\/span><\/a>"
                        }
                        if (u += '<span class="mobile ' + t.Clist[i].M.ClName + '"><a target="_blank" href="' + Domain.WapDomain + '/html/app/appdown.html">' + t.Clist[i].M.Ta + '<\/a><\/span><span class="user-floor">' + t.Clist[i].M.SF + '<\/span><\/div><div class="user-write-msg">' + (t.Clist[i].M.Y !== null && t.Clist[i].M.Y !== "" ? '<span class="user-ip">' + t.Clist[i].M.Y + "<\/span>": "") + '<span class="user-write-time">' + t.Clist[i].M.WT + '<\/span><\/div><div class="user-review">' + t.Clist[i].M.C + '<\/div><div class="review-footer"><span class="review-ft-fr"><span class="stand-by" role="button">支持(' + t.Clist[i].M.S + ')<\/span><span class="oppose" role="button">反对(' + t.Clist[i].M.A + ')<\/span><span class="reply" role="button">回复<\/span><\/span><\/div>', t.Clist[i].R !== null && t.Clist[i].R !== "" && t.Clist[i].R.length > 0) {
                            u += "<ul>";
                            for (r in t.Clist[i].R) if (r <= 3) {
                                u += '<li class="placeholder deputy-floor" data-comment-id="' + t.Clist[i].R[r].Ci + '"><div class="user"><a title="软媒通行证数字ID：' + t.Clist[i].R[r].Ui + '" href="' + t.Clist[i].R[r].UserIndexUrl + '"><div class="user-hd"><img src="' + t.Clist[i].R[r].HeadImg + '" width="45" height="45" onerror="this.src=\'' + Domain.WapImageDomain + '/user/noavatar.png\'" /><\/div><span class="lv">Lv.' + t.Clist[i].R[r].Ul + '<\/span><\/a><\/div><div class="review-con"><div class="user-mes">';
                                switch (t.Clist[i].R[r].M) {
                                case 1:
                                    u += '<a class="rmpvipblue" title="软媒通行证数字ID：' + t.Clist[i].R[r].Ui + '" href="' + t.Clist[i].R[r].UserIndexUrl + '"><span class="user-name">' + t.Clist[i].R[r].N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_blue.svg" width="16" height="15" /><\/a>';
                                    break;
                                case 2:
                                    u += '<a class="rmpvipred" title="软媒通行证数字ID：' + t.Clist[i].R[r].Ui + '" href="' + t.Clist[i].R[r].UserIndexUrl + '"><span class="user-name">' + t.Clist[i].R[r].N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_itred.svg" width="16" height="15" /><\/a>';
                                    break;
                                case 9:
                                    u += '<a class="rmpvip" title="软媒通行证数字ID：' + t.Clist[i].R[r].Ui + '" href="' + t.Clist[i].R[r].UserIndexUrl + '"><span class="user-name">' + t.Clist[i].R[r].N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_orange.svg" width="16" height="15" /><\/a>';
                                    break;
                                default:
                                    u += '<a title="软媒通行证数字ID：' + t.Clist[i].R[r].Ui + '" href="' + t.Clist[i].R[r].UserIndexUrl + '"><span class="user-name">' + t.Clist[i].R[r].N + "<\/span><\/a>"
                                }
                                u += '<span class="mobile ' + t.Clist[i].R[r].ClName + '"><a target="_blank" href="' + Domain.WapDomain + '/html/app/appdown.html">' + t.Clist[i].R[r].Ta + '<\/a><\/span><span class="user-floor">' + t.Clist[i].R[r].SF + '<\/span><\/div><div class="user-write-msg">' + (t.Clist[i].R[r].Y !== null && t.Clist[i].R[r].Y !== "" ? '<span class="user-ip">' + t.Clist[i].R[r].Y + "<\/span>": "") + '<span class="user-write-time">' + t.Clist[i].R[r].WT + '<\/span><\/div><div class="user-review">' + t.Clist[i].R[r].C + '<\/div><div class="review-footer"><span class="review-ft-fr"><span class="stand-by" role="button">支持(' + t.Clist[i].R[r].S + ')<\/span><span class="oppose" role="button">反对(' + t.Clist[i].R[r].A + ')<\/span><span class="reply" role="button">回复<\/span><\/span><\/div><\/div><\/li>'
                            } else {
                                u += '<li class="placeholder deputy-floor hide"  data-comment-id="' + t.Clist[i].R[r].Ci + '"><div class="user"><a title="软媒通行证数字ID：' + t.Clist[i].R[r].Ui + '" href="' + t.Clist[i].R[r].UserIndexUrl + '"><div class="user-hd"><img src="' + t.Clist[i].R[r].HeadImg + '" width="45" height="45"onerror="this.src=\'' + Domain.WapImageDomain + '/user/noavatar.png\'" /><\/div><span class="lv">Lv.' + t.Clist[i].R[r].Ul + '<\/span><\/a><\/div><div class="review-con"><div class="user-mes">';
                                switch (t.Clist[i].R[r].M) {
                                case 1:
                                    u += '<a class="rmpvipblue" title="软媒通行证数字ID：' + t.Clist[i].R[r].Ui + '" href="' + t.Clist[i].R[r].UserIndexUrl + '"><span class="user-name">' + t.Clist[i].R[r].N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_blue.svg" width="16" height="15" /><\/a>';
                                    break;
                                case 2:
                                    u += '<a class="rmpvipred" title="软媒通行证数字ID：' + t.Clist[i].R[r].Ui + '" href="' + t.Clist[i].R[r].UserIndexUrl + '"><span class="user-name">' + t.Clist[i].R[r].N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_itred.svg" width="16" height="15" /><\/a>';
                                    break;
                                case 9:
                                    u += '<a class="rmpvip" title="软媒通行证数字ID：' + t.Clist[i].R[r].Ui + '" href="' + t.Clist[i].R[r].UserIndexUrl + '"><span class="user-name">' + t.Clist[i].R[r].N + '<\/span><img src="//img.ithome.com/images/vip/rmpvip_icon_orange.svg" width="16" height="15" /><\/a>';
                                    break;
                                default:
                                    u += '<a title="软媒通行证数字ID：' + t.Clist[i].R[r].Ui + '" href="' + t.Clist[i].R[r].UserIndexUrl + '"><span class="user-name">' + t.Clist[i].R[r].N + "<\/span><\/a>"
                                }
                                u += '<span class="mobile ' + t.Clist[i].R[r].ClName + '"><a target="_blank" href="' + Domain.WapDomain + '/html/app/appdown.html">' + t.Clist[i].R[r].Ta + '<\/a><\/span><span class="user-floor">' + t.Clist[i].R[r].SF + '<\/span><\/div><div class="user-write-msg">' + (t.Clist[i].R[r].Y !== null && t.Clist[i].R[r].Y !== "" ? '<span class="user-ip">' + t.Clist[i].R[r].Y + "<\/span>": "") + '<span class="user-write-time">' + t.Clist[i].R[r].WT + '<\/span><\/div><div class="user-review">' + t.Clist[i].R[r].C + '<\/div><div class="review-footer"><span class="review-ft-fr"><span class="stand-by" role="button">支持(' + t.Clist[i].R[r].S + ')<\/span><span class="oppose" role="button">反对(' + t.Clist[i].R[r].A + ')<\/span><span class="reply" role="button">回复<\/span><\/span><\/div><\/div><\/li>'
                            }
                            u += "<\/ul>";
                            t.Clist[i].R.length > 4 && (u += '<span class="look-more" role="button">查看更多' + (t.Clist[i].R.length - 4) + "条评论<\/span>")
                        }
                        u += "<\/div><\/li>";
                        $(".comment .all-comment-box .all-comment").append(u)
                    }
                } else $(".comment .all-comment-box .no-comment").removeClass("hide");
                noLoadComment = !1;
                CommentVoteView()
            } else n && $(".comment .all-comment-box .no-comment").removeClass("hide")
        },
        error: function() {
            noLoadComment = !1;
            $(".comment .all-comment-box .comment-tig").remove();
            $(".comment .all-comment-box").append('<div class="comment-tig"><span class="comment-more" role="button">查看更多评论<\/span><\/div>')
        }
    })
}
function NewsGradeShow(n) {
    var t = !0;
    if (n.Grade !== null) switch (n.Me) {
    case 0:
        $(".grade .grade-assess .grade-assess-item").eq(0).addClass("grade-selected");
        t = !1;
        break;
    case 1:
        $(".grade.grade-assess .grade-assess-item").eq(1).addClass("grade-selected");
        t = !1;
        break;
    case 2:
        $(".grade .grade-assess .grade-assess-item").eq(2).addClass("grade-selected");
        t = !1
    }
    t ? ($(".grade .grade-hide-data").show(), $(".grade .grade-data").addClass("hide"), $(".grade .grade-assess .grade-assess-item").removeClass("grade-selected"), $(".grade .grade-assess .grade-assess-item").eq(0).find("span").text("无价值").attr("aria-label", ""), $(".grade .grade-assess .grade-assess-item").eq(1).find("span").text("还可以").attr("aria-label", ""), $(".grade .grade-assess .grade-assess-item").eq(2).find("span").text("有价值").attr("aria-label", ""), n.Grade !== null ? $(".grade .grade-hide-data").text("打分后显示文章质量得分，当前" + (n.G1 + n.G2 + n.G3) + "人打分") : $(".grade .grade-hide-data").text("打分后显示文章质量得分，当前" + ($(".grade .grade-data .grade-num").text() - 1) + "人打分")) : (n.G1 + n.G2 + n.G3 < 5 ? ($(".grade .grade-hide-data").show(), $(".grade .grade-data").addClass("hide"), $(".grade .grade-hide-data").text("打分人数少于5人,不显示文章质量得分")) : ($(".grade .grade-hide-data").hide(), $(".grade .grade-data").removeClass("hide")), $(".grade .grade-data .news-grade").text(n.Grade), $(".grade .grade-data .grade-num").text(n.G1 + n.G2 + n.G3), $(".grade .grade-data .grade-star .grade-star-light").css("width", n.Grade * 20 + "%"), n.Me !== -1 && ($(".grade .grade-assess .grade-assess-item").eq(0).find("span").text(n.G1).attr("aria-label", "无价值"), $(".grade .grade-assess .grade-assess-item").eq(1).find("span").text(n.G2).attr("aria-label", "还可以"), $(".grade .grade-assess .grade-assess-item").eq(2).find("span").text(n.G3).attr("aria-label", "有价值")))
}
function NewsCommentCountGet(n, t) {
    if (t !== "") $.ajax({
        type: "get",
        url: Api.Comment.LapinCommentCountGet,
        data: {
            LapinID: t
        },
        dataType: "json",
        success: function(n) {
            if (n.Success === 1) {
                var t = parseInt(n.Result);
                t >= 1e4 && (t = (t / 1e4).toFixed(2) + "W");
                $(".footer .ft .comment-num .comment-number").text(t)
            } else $(".footer .ft .comment-num .comment-number").text(0)
        }
    });
    else {
        var i = window.sessionStorage.getItem("NewsCommentCount_" + n);
        i === "" || i === null || $.type(i) === "undefined" ? n > 3e5 ? $.ajax({
            type: "get",
            url: Api.Comment.NewsCommentCountGet,
            data: {
                NewsID: n
            },
            dataType: "json",
            success: function(n) {
                if (n.Success === 1) {
                    i = n.Result; (i === "" || i === null || $.type(i) === "undefined" || i === 0 || i === "0") && (i = $(".footer .ft .comment-num .comment-number").attr("data-news-comment-count"), i = parseInt(i));
                    i >= 1e4 && (i = (i / 1e4).toFixed(2) + "W");
                    $(".footer .ft .comment-num .comment-number").text(i);
                    return
                }
            }
        }) : (i === "" || i === null || $.type(i) === "undefined" || i === 0 || i === "0") && (i = $(".comment-number").attr("data-news-comment-count"), i = parseInt(i)) : i = parseInt(i);
        i >= 1e4 && (i = (i / 1e4).toFixed(2) + "W");
        $(".footer .ft .comment-num .comment-number").text(i)
    }
}
var newsID = $("#newsID").attr("data-news-id"),
lapinID = "",
maxCommentID = 0,
latest = "",
userCollectNewsID = localStorage.getItem("UserCollectNewsID"); (userCollectNewsID === "" || $.type(userCollectNewsID) === "undefined" || userCollectNewsID === null) && (userCollectNewsID = "");
var noLoadComment = !1,
upflag = !0,
downflag = !0,
firstLoadComment = !0;
SlideEvent = function(n, t) {
    n === "down" ? (document.documentElement.scrollTop || document.body.scrollTop) >= 78 && $(".fixed-btn").show() : n === "up" && (upflag, (document.documentElement.scrollTop || document.body.scrollTop) <= 78 ? $(".fixed-btn").hide() : $(".fixed-btn").show());
    t === "load" && (firstLoadComment && $(".comment .all-comment-box .all-comment").html() === "" ? AjaxNewsCommonListGet(!0) : AjaxNewsCommonListGet(!1))
};
$(function() {
    var n, t, i;
    if ((matchMedia("(display-mode: standalone)").matches || "standalone" in navigator && navigator.standalone) && $(".footer .ft").prepend('<a aria-label="返回" href="/"><div class="ft-item return-btn"><button><\/button><\/div><\/a>'), sessionStorage.setItem("NewsTitle_" + newsID, $(".con-box .news h1").text()), $(".reward").length > 0 && $.ajax({
        type: "get",
        url: Api.News.NewsRewardGet,
        data: {
            NewsID: newsID
        },
        dataType: "json",
        success: function(n) {
            var i, r, t;
            if (n.Success === 1) {
                for (n = n.Result, i = '<div class="reward-case"><span style="color:#d22222">' + n.Count + '<\/span>人激励原创<\/div><div class="reward-head-img">', r = 0, t = 0; t < (n.Users.length <= 55 ? n.Users.length: 54); t++) i += '<div class="head-img"><img src="' + n.Users[t].UserHeadImg + '" onerror="this.src=\'' + Domain.WapImageDomain + "/user/noavatar.png'\"><\/div>",
                n.Users[t].Money !== null && n.Users[t].Money !== "" && n.Users[t].Money !== undefined && (r += n.Users[t].Money);
                for (n.Count > 55 && (i += '<div class="head-img more-reward-user"><img src="' + Domain.WapImageDomain + '/news/reward-more.png"><\/div>'), i += '<\/div><div class="sdk-mask sdk-mask-close hide"><\/div><div class="reward-collapsed sdk-mask-close hide"><span class="rc-title">' + n.Count + "人打赏" + (r !== 0 ? '<span class="all-money">共' + r.toFixed(2) + "<\/span>": "") + "<\/span><ul>", t = 0; t < n.Users.length; t++) i += '<li><a title="软媒通行证数字ID：' + n.Users[t].UserId + '" href="' + n.Users[t].UserIndexUrl + '"><img src="' + n.Users[t].UserHeadImg + '" onerror="this.src=\'' + Domain.WapImageDomain + '/user/noavatar.png\'"><p class="rc-info"><span class="re-user-nick">' + n.Users[t].UserNick + '<\/span><span class="re-time">' + n.Users[t].AddTimeStr + "<\/span><\/p>" + (n.Users[t].Money !== null && n.Users[t].Money !== "" && n.Users[t].Money !== undefined ? '<p class="rc-money">+' + n.Users[t].Money.toFixed(2) + "<\/p>": "") + "<\/a><\/li>";
                i += "<\/ul><button>收起<\/button><\/div>";
                $(".news .reward").append(i)
            }
        }
    }), $(".click_to_copy_btn").each(function() {
        $(this).removeAttr("onclick").attr("data-clipboard-text", $(this).text())
    }), $.ajax({
        type: "get",
        url: Api.News.NewsGradeGet,
        data: {
            newsID: newsID
        },
        dataType: "json",
        success: function(n) {
            n.Success === 1 && NewsGradeShow(n.Result)
        }
    }), $.ajax({
        type: "get",
        url: Api.LaPin.OurBuyGet,
        dataType: "json",
        success: function(n) {
            var t, i, r;
            if (n.Success === 1) {
                $(".con-box .lapin").removeClass("hide");
                for (t in n.Result) i = '<div class="placeholder swiper-slide"><a href="' + (n.Result[t].QuanUrl === null && n.Result[t].QuanUrl === "" && $.type(n.Result[t].QuanUrl) === "undefined" ? n.Result[t].BuyUrl: n.Result[t].QuanUrl) + '"><div class="plc-image"><img class="swiper-lazy" data-src="' + n.Result[t].Picture_square + '" /><\/div><div class="plc-con"><p class="plc-title">' + n.Result[t].ProductName + '<\/p><div class="price"><div class="coupon"><span class="rel-price">' + n.Result[t].PromotionInfoPrice + '<\/span><span class="get-coupon ' + (n.Result[t].QuanUrl === null && n.Result[t].QuanUrl === "" && n.Result[t].QuanUrl === undefined ? "go-buy": "get-quan") + '" role="button">' + n.Result[t].QuanInfo + "<\/span><\/div><\/div><\/div><\/a><\/div>",
                $(".con-box .lapin .lapin-box.swiper-container .swiper-wrapper").append(i);
                r = new Swiper(".lapin .lapin-box.swiper-container", {
                    autoplay: !1,
                    uniqueNavElements: !1,
                    freeMode: !0,
                    preloadImages: !0,
                    slidesPerView: "auto",
                    grabCursor: !0,
                    spaceBetween: 0,
                    roundLengths: !0,
                    lazy: {
                        loadPrevNext: !0,
                        loadPrevNextAmount: 4,
                        elementClass: "swiper-lazy"
                    }
                })
            }
        }
    }), userInfo !== null && userInfo.NoComment) noLoadComment = !0,
    $(".comment,.footer .comment-num").remove(),
    $(".footer .write-review span").text("您已关闭评论系统"),
    $(".footer .write-review").css({
        "pointer-events": "none"
    });
    else {
        NewsCommentCountGet(newsID, lapinID);
        $(document).on("click", ".comment .no-comment-box .no-comment",
        function() {
            parentCommentID = "";
            ppcID = "";
            replyTo = "";
            OpenComment()
        })
    }
    ScrollDirect(SlideEvent);
    window.location.href.match("#mao-comment") && $("html,body").stop(!0, !0).animate({
        scrollTop: $(".con-box .comment").offset().top
    },
    370,
    function() {
        $("html,body").scrollTop($(".con-box .comment").offset().top)
    });
    $(document).on("click", ".down-app-box .down-app,.open-app-a",
    function(n) {
        OpenIthomeAppByA("news", "id=" + newsID, n)
    });
    userCollectNewsID.indexOf("|" + newsID) !== -1 && ($(".footer .footer-box .ft .news-collect button").attr("aria-checked", !0), $(".footer .footer-box .ft .news-collect").removeClass("news-collect-gray").addClass("news-collect-red"));
    $(document).on("click", ".comment .all-comment-box .comment-tig .comment-more",
    function() {
        AjaxNewsCommonListGet(!1)
    });
    $(document).on("click", ".footer .footer-box .ft .ft-item.comment-num",
    function() {
        $("html,body").stop(!0, !0).animate({
            scrollTop: $(".con-box .comment").offset().top
        },
        370,
        function() {
            $("html,body").scrollTop($(".con-box .comment").offset().top)
        })
    });
    n = !1;
    $(document).on("click", ".footer .footer-box .ft .news-collect",
    function() {
        if (n) return ! 1;
        n = !0;
        var i = $(this),
        t = "add";
        i.hasClass("news-collect-gray") ? t = "add": i.hasClass("news-collect-red") && (t = "del");
        $.ajax({
            type: "post",
            url: Api.News.UserNewsCollectSet,
            data: {
                newsID: newsID,
                type: t
            },
            dataType: "json",
            success: function(r) {
                r.Success === 1 ? (userCollectNewsID = userCollectNewsID.replace("|" + newsID, ""), t === "add" ? (userCollectNewsID += "|" + newsID, localStorage.setItem("UserCollectNewsID", userCollectNewsID), i.removeClass("news-collect-gray").addClass("news-collect-red")) : t === "del" && (localStorage.setItem("UserCollectNewsID", userCollectNewsID), i.removeClass("news-collect-red").addClass("news-collect-gray")), n = !1) : (r.Result === "needLogin" && LoginWindow(), n = !1)
            }
        })
    });
    $(document).on("click", ".footer .footer-box .ft .news-share",
    function() {
        var n = $(".news-content img:first").attr("src"); (n === "" || n === null || $.type(n) === "undefined") && (n = "");
        NShareSet(n, window.location.href, $("title").text(), $("meta[name=description]").attr("content"), "@IT之家");
        NShareCall()
    });
    $(document).on("click", ".reward .reward-head-img",
    function() {
        $(".reward .sdk-mask,.reward .reward-collapsed").removeClass("hide")
    });
    $(document).on("click", ".reward .reward-collapsed button",
    function() {
        $(".reward .sdk-mask,.reward .reward-collapsed").addClass("hide")
    });
    t = !1;
    $(document).on("click", ".grade .grade-assess .grade-assess-item",
    function() {
        if (t) return ! 1;
        t = !0;
        var n = $(this);
        n.hasClass("grade-selected") ? $.ajax({
            type: "post",
            url: Api.News.NewsGradeCannel,
            data: {
                newsID: newsID
            },
            success: function(n) {
                n.Success === 1 ? NewsGradeShow(n.Result) : n.Result === "needLogin" ? LoginWindow() : n.Result !== "" && typeof n.Result != "undefined" ? layer.msg(n.Result) : layer.msg("取消评分失败");
                t = !1
            }
        }) : $.ajax({
            type: "post",
            url: Api.News.NewsGradeSet,
            data: {
                newsID: newsID,
                grade: n.index()
            },
            success: function(n) {
                n.Success === 1 ? NewsGradeShow(n.Result) : n.Result === "needLogin" ? LoginWindow() : n.Result !== "" && n.Result !== null && typeof n.Result != "undefined" ? layer.msg(n.Result) : layer.msg("评分失败");
                t = !1
            }
        })
    });
    $(document).on("click", ".comment .re-tags .tags-sort label",
    function() {
        latest = $(this).attr("data-latest");
        var n = $(".comment .re-tags .tags-sort input[type='radio']:checked + label").attr("data-latest");
        n !== latest && (maxCommentID = latest === "1" ? 1 : 0, $(".comment .all-comment").html(""), noLoad = !1, noLoadComment = !1, AjaxNewsCommonListGet(!1))
    });
    i = !1;
    $(document).on("click", ".relevant-news .more-relevant-news",
    function() {
        if (i) return ! 1;
        i = !0;
        var n = $(this),
        t = $(".relevant-news .placeholder:last").attr("data-order-newsId");
        $.ajax({
            type: "get",
            url: Api.News.MoreRelatedNewsGet,
            data: {
                newsId: newsID,
                orderNewsId: t
            },
            success: function(t) {
                var r, u;
                if (t.Success === 1) for (t.Result.bnp && t.Result.NewsList.length !== 0 || n.remove(), t = t.Result.NewsList, r = 0; r < t.length; r++) u = '<div class="placeholder one-img-plc" data-order-newsId="' + t[r].newsid + '"><a href="javascript:window.location.replace(\'' + t[r].WapUrl + '\');"><div class="plc-image"><img src="' + t[r].img + '" /><\/div><div class="plc-con"><p class="plc-title">' + t[r].newstitle + '<\/p><p class="plc-footer plc-footer-fr"><span>' + t[r].postdate + "<\/span><\/p><\/div><\/a><\/div>",
                $(".relevant-news-box").append(u);
                else n.remove();
                i = !1
            },
            error: function() {
                i = !1
            }
        })
    });
    $(document).on("click", ".comment ul .placeholder .look-more",
    function() {
        $(this).hide();
        $(this).parent().find("ul li").removeClass("hide")
    });
    $("a").each(function() {
        var n = $(this).attr("data-wapurl");
        n !== undefined && n !== "" && n !== null && $(this).attr("href", n)
    });
    $(document).on("click", ".news-content img",
    function() {
        var t = 0,
        r = $(this).attr("data-original"),
        n = '<div class="img-box swiper-container"><div class="swiper-wrapper">',
        i;
        $(".news-content img").each(function(i) {
            $(this).attr("data-original") === r && (t = i);
            n += '<div class="swiper-slide"><div class="swiper-zoom-container"><img class="swiper-lazy" data-src="' + $(this).attr("data-original") + '"><\/div><\/div>'
        });
        n += '<\/div><div class="swiper-pagination"><\/div><\/div>';
        $("body").append(n);
        i = new Swiper(".img-box.swiper-container", {
            autoplay: !1,
            preloadImages: !0,
            slidesPerView: 1,
            grabCursor: !0,
            spaceBetween: 0,
            roundLengths: !0,
            lazy: {
                loadPrevNext: !0,
                loadPrevNextAmount: 3,
                elementClass: "swiper-lazy"
            },
            pagination: {
                el: ".img-box.swiper-container .swiper-pagination",
                type: "progressbar",
                clickable: !0
            },
            zoom: !0
        });
        i.slideTo(t, 0, !1)
    });
    $(document).on("click", ".img-box",
    function() {
        $(".img-box").remove()
    })
});