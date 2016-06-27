import JSONSerializer from 'ember-data/serializers/json';

export default JSONSerializer.extend({
});

// import DS from 'ember-data';
//
// export default DS.RESTSerializer.extend({
//   serialize: function(snapshot, options) {
// 	var json = {
// 	  POST_TTL: snapshot.attr('title'),
// 	  POST_BDY: snapshot.attr('body'),
// 	  POST_CMS: snapshot.hasMany('comments', { ids: true })
// 	}
//
// 	if (options.includeId) {
// 	  json.POST_ID_ = snapshot.id;
// 	}
//
// 	return json;
//   }
// });
