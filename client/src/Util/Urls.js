const Urls = {};

if (process.env.NODE_ENV === 'local') {
  Urls.api = 'http://localhost:5000/api';
} else {
  Urls.api = `${process.env.PUBLIC_URL}/api`; // can be different than Dev if needed
}

export default Urls;