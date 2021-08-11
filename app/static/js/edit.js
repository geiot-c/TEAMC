
//暇ボタン
var shops = new Vue({
    el:'#app',
    //{{}}を[()]に変更
    delimiters: ['[(', ')]'],
    data:{
        message: "hello world!",
        selcted:null
    },
    mounted () 
    {
        this.selcted=Number(location.pathname.split('/')[2]);
        
    },
    methods: {
        send: function() {
            console.log("send");
            $.post(location.href+"/is_hot",{
                'is_hot': true
              })
            .done(function( data ) {
                console.log("ooo");
            });
        }
    }
});