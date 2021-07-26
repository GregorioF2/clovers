const axios = require("axios");
const { SERVER_URL } = require("../configs");

exports.obtainJugRiddleSolution = async (x, y, z) => {
  try {
    const res = await axios.get(`${SERVER_URL}/riddles/jug`, {
      params: {
        x,
        y,
        z,
      },
    });
    return res.data;
  } catch (err) {
    throw err.response.data;
  }
};
