export default {
  install: (app, options) => {
    console.log("----------------------------------------", app, options);
    app.provide("formatNanoseconds", formatNanoseconds);
    app.config.globalProperties.$formatNanoseconds = formatNanoseconds;
  },
};

function formatNanoseconds(nanoseconds) {
  const totalSeconds = nanoseconds / 1e9;
  const hours = Math.floor(totalSeconds / 3600);
  const minutes = Math.floor((totalSeconds % 3600) / 60);
  const seconds = Math.floor(totalSeconds % 60);

  let formattedString = "";

  if (hours > 0) {
    formattedString += `${hours}h`;
  }

  if (minutes > 0 || hours > 0) {
    formattedString += `${minutes}m`;
  }

  formattedString += `${seconds}s`;

  return formattedString;
}
