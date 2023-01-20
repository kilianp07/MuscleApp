const firebaseConfig = require("../../config/firebase");
const firebaseAuth = require("firebase/auth");
const firebaseApp = require("firebase/app");
const app = firebaseApp.initializeApp(firebaseConfig.config);
const auth = firebaseAuth.getAuth();

function incomingUserCorrectlyFilled(user)  {
    return user.email !=null && user.password !=null && user.username !=null;
}

async function createUserWithEmailAndPassword(email, password) {
  try{
    const userCredential = await firebaseAuth.createUserWithEmailAndPassword(auth, email, password)
    return userCredential.user;
  }
  catch(error) {
    throw new Error(error.code);
  }
}

async function signInWithEmailAndPassword(email, password) {
  try{
    const userCredential = await firebaseAuth.signInWithEmailAndPassword(auth, email, password)
    return userCredential.user;
  }
  catch(error) {
    throw new Error(error.code);
  }
}

// handleRegister
async function register (username, email, password) {
  try {
    await firebaseAuth.createUserWithEmailAndPassword(auth, email, password);
    await firebaseAuth.sendEmailVerification(auth.currentUser)
    await firebaseAuth.updateProfile(auth.currentUser, { displayName: username })
    return auth.currentUser
  } catch (err) {
    throw new Error(err.code)
  }
};

async function logout () {
  try {
    await firebaseAuth.signOut(auth)
  } catch (err) {
    throw new Error(err.code)
  }
  return true
}

module.exports = {
    createUserWithEmailAndPassword,
    incomingUserCorrectlyFilled,
    signInWithEmailAndPassword,
    register,
    logout
}
