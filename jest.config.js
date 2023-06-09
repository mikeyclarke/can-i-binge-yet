module.exports = {
    "roots": [
        "<rootDir>/tests/ts"
    ],
    transform: {
        "\\.(ts|js)$": "ts-jest",
    },
    modulePaths: ['<rootDir>/src/ts'],
    testEnvironment: "jsdom",
    setupFiles: ['<rootDir>/tests/ts/setupJest.ts'],
}
