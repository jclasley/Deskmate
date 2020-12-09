const Urls = {};

if (process.env.NODE_ENV === 'production') {
  Urls.api = 'http://localhost:5000/api'; // can be different than Dev if needed
} else if (process.env.NODE_ENV === 'development') {
  Urls.api = 'http://localhost:5000/api';
}

export default Urls;