// Get latest version

const latestURL = 'https://api.github.com/repos/rytsh/indigo/releases/latest';

const getData = async (url = latestURL) => {
    // Default options are marked with *
    const response = await fetch(url);
    return await response.json(); // parses JSON response into native JavaScript objects
};

export { getData };
