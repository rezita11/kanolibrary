<template>
    <div>

        <pre>{{ book }}</pre>

        <v-switch v-model="book.Available">

        </v-switch>
        <v-text-field title="Name" v-model="book.Name" :rules="[v => !!v || 'Name is required']"></v-text-field>
        <v-text-field title="Author" v-model="book.Author" :rules="[v => !!v || 'Author is required']"></v-text-field>
        <v-text-field title="Publisher" v-model="book.Publisher" :rules="[v => !!v || 'Publisher is required']"></v-text-field>
        <v-text-field title="Synopsis" v-model="book.Synopsis" :rules="[v => !!v || 'Synopsis is required']"></v-text-field>
        <v-text-field title="Page" v-model="book.Page" :rules="[v => !!v || 'Page is required']"></v-text-field>
        <v-btn @click="save">
            Save
        </v-btn>
        <v-btn @click="close">
            Close
        </v-btn>
    </div>
</template>
<script>
export default {
    name: 'Books-id',
    data() {
        return {
            bookId: this.$route.params.id,
            book: {},
            isNew: false,
        }
    },
    methods: {
        save() {
            if (this.isNew) {
                this.$axios.$post('/api/books/insert', this.book)
                    .then(response => {
                        this.$router.push('/Books')
                    })
            } else {
                this.$axios.$post('/api/books/update?id=' + this.bookId, this.book)
                    .then(response => {
                        this.$router.push({
                            path: '/Books',
                        })
                    })
            }

        },
        close() {
            this.$router.push({
                path: '/Books',
            })
        }
    },
    mounted() {
        // this.$axios.$get('/api/books/' + this.bookId)
        //     .then(response => {
        //         this.book = response
        //         console.log(response)
        //     })
        if (this.bookId !== 'add') {
            this.$axios.$get('/api/books/find?id=' + this.bookId)
                .then(response => {
                    if (response.length > 0) {
                        this.book = response[0]
                    } else {
                        this.$router.push('/Books')
                    }
                    // this.book = response
                    // console.log(response)
                })
        }
        else {
            this.isNew = true
        }
    },
}

</script>
