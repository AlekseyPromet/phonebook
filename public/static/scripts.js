var loc = window.location;
var uri = 'ws://localhost:8080/search';
var ws = new WebSocket(uri);

var v1 = new Vue({
  el: '#app1',
  data: {
    searchMessage:
      'Начните набирать номер телефона безы +7, например 987654321',
    allNumbers: [],
    searchNumber: '',
    isSearch: false,
  },
  created() {
    ws.onopen = function(v1) {
      console.log('Подключаемся к вёбсокету');
    };
  },
  mounted() {
    console.log(uri);
  },
  destroyed() {
    ws.onclose = function() {
      if (ws.wasClean) {
        alert('Соединение закрыто чисто');
      } else {
        // например, "убит" процесс сервера
        alert('Обрыв соединения');
      }
      alert('Код: ' + ws.code + ' причина: ' + ws.reason);
    };
  },
  methods: {
    search: function() {
      console.log('Отправляем данные на сервер');
      setInterval(function() {
        ws.send(this.searchNumber);
      }, 10000);
      ws.onmessage = function() {
        console.log('Ответ от сервера', ws.data);
        this.allNumbers += ws.data;
        console.log(v1.allNumbers);
        this.isSearch = true;
      };
    },
  },
});

var v2 = new Vue({
  el: '#app2',
  data: {
    contacts: [
      {
        id: 0,
        firstname: '',
        secondname: '',
        sinonim: '',
        prefix: '',
        number: 0,
        active: true,
      },
    ],
    isLoading: false,
    locContacts: {},
  },
  //Получаем данные при создании страницы
  created() {
    this.getData();
  },
  //Получаем данные при обновлении страницы
  methods: {
    getData: function() {
      return axios
        .get('/contacts')
        .then(function(response) {
          v2.contacts = response.data.contacts ? response.data.contacts : {};
          v2.isLoading = true;
        })
        .catch(function(error) {
          console.log('Ошибка при запроске контактов');
        });
    },
    //Delete contact from db
    deleteContact: function(index) {
      axios
        .delete('/delсontact/' + v2.contacts.id)
        .then(response => {
          v2.contacts.splice(index, 1);
          console.log('Контакт id=' + index + ' удалён');
        })
        .catch(error => {
          console.log('Не удалось удалить контакт id=' + index);
        });
    },
  },
});

var v3 = new Vue({
  el: '#app3',
  data: {
    newContact: {
      firstname: '',
      secondname: '',
      sinonim: '',
      prefix: '',
      number: null,
      active: true,
    },
  },
  methods: {
    //Создать котакт
    createContact: function() {
      //Если
      if (!v3.newContact.number) {
        console.log('Пустой номер');
        this.$set(v3.newContact.fisrtname, '');
        this.$set(v3.newContact.secondname, '');
        this.$set(v3.newContact.sinonim, '');
      } else {
        axios
          .put('/newсontact', v3.newContact)
          .then(function(response) {
            this.newContact.id = response.created.id;
            this.contacts.push(v3.newContact);
            console.log('Контакт создан!');
          })
          .catch(function(error) {
            console.log('Ошибка создания нового контакта');
          });
      }
    },
  },
});
