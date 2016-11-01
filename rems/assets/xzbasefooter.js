
//点击左侧菜单 
$('#proj li').click(function() {
    //收缩所有
    $('#proj li').removeClass("active");
    
    // 找出 li 中的超链接 href(#id)  跳转 iframe中的内容
    var $this = $(this);
    $this.addClass("active");
    _clickTab = $this.find('a').attr('atarget'); // 找到链接a中的targer的值
    var $jd = {"src":_clickTab};
    $("#myiframe").prop($jd);    
    
    $.get(_clickTab,function(data){
    //$("#iframe").html(data); 
    });
});
   

        
//头部下拉导航      
$(".flip").click(function(){
    $(".newnavi").slideToggle("fast");
  });
$(".newnavi").mouseleave(function(){
   $(".newnavi").slideToggle("fast");
}) 
 
         
         
 