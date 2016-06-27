import Ember from 'ember';

export default Ember.Route.extend({
    model() {
        let todos = this.store.findAll("todo");
        return Ember.RSVP.hash({
            todos: todos
        });
    },

    setupController(controller, models) {
        controller.set("todos", models.todos);
    },

    actions: {
        addNewTodo(todoName) {
            let todo = this.store.createRecord("todo", {
                name: todoName
            });
            todo.save();
        },

        deleteTodo(todo) {
            this.store.findRecord('todo', todo.id).then(function(t) {
                t.destroyRecord(); // => DELETE to /posts/2
            });
        },

        updateTodo(todoId, todoName) {
            this.store.findRecord('todo', todoId).then(function(t) {
                t.set("name", todoName);
                t.save();
            });
        }
    }

});
