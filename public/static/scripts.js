var v1 = new Vue({
  el: '#app1',
  data: {
    searchMessage:
      'Начните набирать номер телефона безы +7, например 987654321',
    numbers: [],
    searchNumber: '',
  },
  methods: {
    search: function(value) {
      this.searchNum = toString(value)
        .split('-')
        .join('');
      return axios
        .get('/search/' + this.searchNum)
        .then(response => {
          this.numbers = response.data.number ? response.data.number : [];
        })
        .catch(error => {
          console.log('Ошибка при запросе номеров');
        });
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
          .then(response => {
            this.newContact.id = response.created;
            this.contacts.push(v3.newContact);
            console.log('Контакт создан!');
          })
          .catch(error => {
            console.log('Ошибка создания нового контакта');
          });
      }
    },
  },
});
