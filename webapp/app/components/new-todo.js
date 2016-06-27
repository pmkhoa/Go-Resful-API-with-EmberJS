import Ember from 'ember';

export default Ember.Component.extend({
    isAddingTodo: false,
    isEmptyTodoName: false,
    todoName: "",

    actions: {
        addNewTodo() {
            this.set("isAddingTodo", true);
        },
        saveTodo() {
            let todoName = this.get("todoName");
            if ( todoName !== "" ) {
                this.sendAction("addNewTodo", this.get("todoName"));
                this.set("isAddingTodo", false);
                this.set("isEmptyTodoName", false); 
                this.set("todoName", "");
            } else {
                this.set("isEmptyTodoName", true); 
            }
        }
    }
});
