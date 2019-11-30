<template>
    <Page>
        <ActionBar>
            <Label text="Beer Store"></Label>
        </ActionBar>

        <GridLayout columns="*,auto">
            <TextField hint="Search beer" v-model="textFieldValue" />
        </GridLayout>
    </Page>
</template>

<script>
    export default {
        computed: {
            message() {
                return "Blank {N}-Vue app";
            }
        },
        data(){
            return {
                textFieldValue:""
            }
        }
    };
</script>

<style scoped lang="scss">
    @import '~@nativescript/theme/scss/variables/blue';

    // Custom styles
    .fas {
        @include colorize($color: accent);
    }

    .info {
        font-size: 20;
        horizontal-align: center;
        vertical-align: center;
    }
</style>
