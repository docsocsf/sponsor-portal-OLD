export default [
  {
    pattern: '(.*)',

    fixtures: function (match, params, headers, context) {
      if (match[1] === '/cv') {
        context.progress = {
          parts: 50,
          delay: 40,
        };
        return 'CV uploaded!';
      }
    },
    post: function (match, data) {
      return {
        code: 201
      };
    }
  },
];
