import Ember from 'ember';

export default Ember.Component.extend({
    isUpdatingTodo: false,
    tagName: 'li',
    classNames: ['single-todo'],

    actions: {
        deleteTodo(todo) {
            this.sendAction("deleteTodo", todo);
        },
        addNewTodo() {
            this.sendAction("addNewTodo");
        },
        editTodo() {
            this.toggleProperty("isUpdatingTodo");
        },
        updateTodo(todo) {
            let todoName = this.get("todo.name");
            let todoId = this.get("todo.id");
            if ( todoName !== "" ) {
                this.sendAction("updateTodo", todoId, todoName);
                this.toggleProperty("isUpdatingTodo");
            }
            //TODO: Show error when enter empty
        }
    }
});
