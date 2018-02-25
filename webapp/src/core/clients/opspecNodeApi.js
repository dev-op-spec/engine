import OpspecNodeApiClient from '@opspec/sdk/lib/node/api/client';

const config = {};
if (process.env.NODE_ENV === 'production') {
  // in production build, we assume the node API we talk to is available via current protocol & host.
  // this differs from development build where we talk to local node API
  config.baseUrl = `${window.location.protocol}//${window.location.host}/api`;
} else{
  config.baseUrl = `${window.location.protocol}//localhost:42224/api`
}

export default new OpspecNodeApiClient(config);
