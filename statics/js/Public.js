function RefreshHtml(){location.reload()}function ReLocation(a){if("back"==a)return RefreshHistory();window.location.href=a}function GetDataById(a,b,c){a=document.getElementById(a).value;10<a.length?alert("\u8bf7\u8f93\u516510\u4e2a\u5b57\u7b26\u4ee5\u5185\u7684\u67e5\u8be2\u5b57\u6bb5\uff01\uff01\uff01"):(""==a&&(a=""),window.location="/User/myCourse?id="+b+"&loginCode="+c+"&search_data="+a)}
function delCourseByField(a,b){return 1==confirm("\u662f\u5426\u5bf9["+a+b+"]\u8fdb\u884c\u9000\u8bfe\uff1f")?!0:!1}function executeSql(a,b){1==a&&(window.location.href=b)}function StudAddCourse(a,b,c){confirm("\u786e\u8ba4\u9009\u62e9\u8be5\u8bfe\u7a0b\u5417\uff1f")&&(window.location.href="/User/studAddCourse?id="+a+"&C_id="+b+"&loginCode="+c)}
function StudSearchCourse(a,b,c){a=document.getElementById(a).value;10>=a.length?window.location.href="/User/joinCourse?id="+b+"&loginCode="+c+"&search_words="+a:alert("\u8bf7\u8f93\u5165\u5341\u4e2a\u5b57\u7b26\u4ee5\u5185\u7684\u5173\u952e\u8bcd\uff01")}
function PersonnelManage(a,b,c){2>a?alert("\u60a8\u6ca1\u6709\u6743\u9650\uff01"):(alert("\u6b22\u8fce\u8fdb\u5165\u4eba\u5458\u7ba1\u7406\u754c\u9762\uff0c\u8d85\u7ea7\u7ba1\u7406\u5458\uff01"),window.location="/User/personalManage?id="+b+"&loginCode="+c)}function PersonalSearch(a,b){var c=document.getElementById(a).value;10<c.length?(alert("\u641c\u7d22\u8bcd\u8fc7\u957f\uff0c\u8bf7\u572810\u4e2a\u5b57\u7b26\u4ee5\u5185\uff01"),b.reload(b)):window.location=b+"&search_words="+c}
function StudManage(a,b,c){a=document.getElementById(a).value;b=document.getElementById(b).value;20<a.length?alert("\u641c\u7d22\u8bcd\u8fc7\u957f"):window.location=c+"&searchWord="+a+"&courseId="+b}function DelStudentBySid(a){alert(a);1==confirm("\u7b2c\u4e00\u6b21\u786e\u8ba4\u5220\u9664?")&&confirm("\u7b2c\u4e8c\u6b21\u786e\u8ba4\u5220\u9664?")&&(window.location=a)}function GradesDada(a,b){var c=document.getElementById(a).value;window.location=b+"&score="+c}
function GetEidHref(a,b,c){a=document.getElementById(a).value;window.location=c+"&"+b+"="+a}function HrefHtml(a){window.location=a}function RefreshHistory(){window.location.replace(document.referrer)}function DiyConfirm(a,b){1==confirm(a)&&(window.location=b)}function SaveCourse(a,b,c,d){a=document.getElementById(a).value;b=document.getElementById(b).value;c=document.getElementById(c).value;window.location=d+"&Cname="+a+"&Teacher="+b+"&Tscore="+c}
function AdminResetPasswd(a,b,c,d){a=document.getElementById(a).value;b=document.getElementById(b).value;c=document.getElementById(c).value;window.location=d+"&old_p="+a+"&new_p1="+b+"&new_p2="+c}function PagingHref(a,b,c){b=document.getElementById(b).value;b>c||0>=b?alert("\u9875\u9762\u4e0d\u5b58\u5728\uff01\uff01\uff01"):window.location=a+"&search_Page="+b}
function LastPage(a,b,c){b=document.getElementById(b).value;b=parseInt(b);c=parseInt(c);b>c&&(b=c);0<b-1?(--b,a=a+"&search_Page="+b):(alert("\u5f53\u524d\u662f\u7b2c\u4e00\u9875\uff0c\u4e0d\u80fd\u7ee7\u7eed\u4e0a\u4e00\u9875\u4e86\uff01\uff01\uff01"),a+="&search_Page=1");window.location=a}
function NextPage(a,b,c){b=document.getElementById(b).value;b=parseInt(b);b+1<=c?a=a+"&search_Page="+(b+1):(alert("\u5f53\u524d\u662f\u6700\u540e\u4e00\u9875\uff0c\u4e0d\u80fd\u7ee7\u7eed\u4e0b\u4e00\u9875\u4e86\uff01\uff01\uff01"),a=a+"&search_Page="+c);window.location=a}function StudentLastPage(a,b,c,d){if(""==d||"0"==d)d="ALL";LastPage(a+("&courseId="+d),b,c)}function StudentNextPage(a,b,c,d){if(""==d||"0"==d)d="ALL";NextPage(a+("&courseId="+d),b,c)}
function StudentPagingHref(a,b,c,d){if(""==d||"0"==d)d="ALL";PagingHref(a+("&courseId="+d),b,c)}function isEmptyStr(a){return void 0==a||null==a||""==a?!0:!1}
function JsAddFooter(a,b,c){var d=document.getElementsByTagName("footer")[0];d.style.width=b;d.style.height=a;d.style.backgroundColor="deepskyblue";d.style.position="relative";d.style.border="solid skyblue 6px";d.style.top=c;a=document.createElement("p");b=document.createElement("p");c=document.createElement("p");d=document.createElement("a");a.innerHTML="\u514d\u8d23\u58f0\u660e\uff1a\u672c\u7ad9\u4e3a\u4e2a\u4eba\u7f51\u7ad9\uff0c\u7528\u4e8e\u5b66\u4e60\u548c\u7814\u7a76;\u4e0d\u5f97\u5c06\u4e0a\u8ff0\u5185\u5bb9\u7528\u4e8e\u5546\u4e1a\u6216\u8005\u975e\u6cd5\u7528\u9014\uff01";
b.innerHTML="\u672c\u7ad9\u4e3a\u975e\u76c8\u5229\u6027\u7ad9\u70b9,\u6240\u6709\u5185\u5bb9\u4e0d\u4f5c\u4e3a\u5546\u4e1a\u884c\u4e3a\u3002";c.innerHTML="Copyright \u00a9 2022 \u5b66\u751f\u9009\u8bfe\u7cfb\u7edf--";d.innerHTML="\u9c81ICP\u59072022025031\u53f7-1";d.href="https://beian.miit.gov.cn/";d.target="_blank";d.style.textDecoration="none";c.appendChild(d);d=document.getElementById("JS_Footer");d.appendChild(a);d.appendChild(b);d.appendChild(c);d.style.position="absolute";d.style.left="30%";
d.style.fontSize="20px"}function mOver(a){a.innerHTML="9999";document.getElementById("hid_div").style.display="block"}function mOut(a){document.getElementById("hid_div").style.display="none"}function Test(a){a=document.getElementById(a).value;alert(a)};
