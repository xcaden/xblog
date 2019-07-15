Vue.component('editor',{
    template:'<p>Hey there, I am {{name}}. <button v-on:click="changeName">Change name</button></p>',
    data: function(){
        return {
            name: 'Yoshi'
        }
    },
    methods: {
        changeName: function(){
            this.name = 'Mario';
        }
    }
});

new Vue({
    el: '#app1'
});