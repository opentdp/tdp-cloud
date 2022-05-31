const app = Vue.createApp({});

app.component('app-home', {
    data() {
        return {
            loding: false,
            SecretId: '',
            SecretKey: '',
        };
    },
    watch: {
        message(newValue, oldValue) {
        }
    },
    methods: {
        login() {
            this.loading = true;
            fetch(`api/login?SecretId=${this.SecretId}&SecretKey=${this.SecretKey}`)
                .then(response => response.json())
                .then(data => {
                    const items = [0, 1, 2, 3].map(i => {
                        return data[i] || this.items[i];
                    });
                    this.items = items.filter(v => v);
                    this.pagerRender();
                })
                .catch(err => {
                    this.pullMsg = 2;
                })
                .finally(() => {
                    this.pulling = false;
                });
        }
    },
    template: `
        <nav class="navbar navbar-expand-lg navbar-light bg-light">
            <div class="container-xxl justify-content-start">
                <a class="navbar-brand">
                   <img src="assets/images/tdp-logo.png" style="height: 50px" />
                </a>
                <div class="text-secondary">TDP Cloud</div>
            </div>
        </nav>
        <div class="container-xxl mt-3">
            <div class="mb-3">
                <input type="text" class="form-control" placeholder="SecretId" v-model="SecretId">
            </div>
            <div class="mb-3">
                <input type="text" class="form-control" placeholder="SecretKey" v-model="SecretKey">
            </div>
        </div>
    `
});

app.mount('app-root');
