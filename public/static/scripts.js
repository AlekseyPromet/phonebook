new Vue({
  el: 'body',
  data: {
    numbers: [],
    contacts: [{
      id: '',
      firstname: '',
      secondname: '',
      sinonim: '',
      prefix: '',
      number: '',
      active: '',
    }],
    contact: {
      id: '',
      firstname: '',
      secondname: '',
      sinonim: '',
      prefix: '',
      number: '',
      active: '',
    }
  },
  //Получаем данные при создании страницы
  created() {
    this.getData();
  },
  //Получаем данные при обновлении страницы
  updated() {
    this.getData();
  },
  methods: {
    getData() {
      axios.get('/contacts').then(response => {
        this.contacts = JSON.parse(response.data.contacts) ? response.data.contacts : [];
      }).catch(e => {
        console.log('Ошибка при запроске контактов');
      });
    },
    search() {
      axios.get('/numbers').then(response => {
        this.numbers = response.data.numbers ? response.data.numbers : [];
      }).catch(e => {
        console.log('Ошибка при запроске номеров');
      });
    },
    //Создать котакт
    createContact() {
      //Если
      if (!this.newContact.number.trim()) {
        console.log('Пустой номер');
        this.$set(this.newContact, 'fisrtname', "");
        this.$set(this.newContact, 'secondname', "");
        this.$set(this.newContact, 'sinonim', "");
        return;
      } else {
        this.$set(this.newContact, 'firstname', this.newContact.fisrtname.trim());
        this.$set(this.newContact, 'secondname', this.newContact.secondname.trim());
        this.$set(this.newContact, 'sinonim', this.newContact.sinonim.trim());
        this.$set(this.newContact, 'prefix', this.newContact.prefix);
        this.$set(this.newContact, 'number', parseInt(this.newContact.fisrtname.trim()));
        this.$set(this.newContact, 'active', this.newContact.active);
        axios.put('/newсontact', this.newContact)
          .then(response => {
            this.newContact.id = response.created;
            this.contacts.push(this.newContact);
            console.log("Контакт создан!");
          })
          .then(() => {
            this.$set(this.newContact, 'id', "");
            this.$set(this.newContact, 'firstname', "");
            this.$set(this.newContact, 'secondname', "");
            this.$set(this.newContact, 'sinonim', "");
            this.$set(this.newContact, 'prefix', "7");
            this.$set(this.newContact, 'number', "");
            this.$set(this.newContact, 'active', "true");
          })
          .catch(error => {
            console.log('Ошибка создания нового контакта');
          });
      }
    },
    //Delete task from db
    deleteContact(index) {
      axios.delete('/delсontact/' + this.contacts[index].id)
        .then(response => {
          this.tasks.splice(index, 1);
          console.log("Контакт id=" + index + " удалён");
        })
        .catch(error => {
          console.log('Не удалось удалить контакт id=' + index);
        });
    }
  },
});