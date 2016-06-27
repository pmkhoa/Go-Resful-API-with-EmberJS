import Ember from 'ember';

export default Ember.Component.extend({
    uncomppletedTodos: Ember.computed("todos.@each", function() {
        console.log(this.get("todos"));
        return this.get("todos.length");
    })
});
