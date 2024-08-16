// import { useContext } from "react";
//import Signin from "./pages/Signin";

import { appRouterSwitch } from "./routers";
// import Message from "./components/Message"

function App() {
  // const {

  // } = useContext
  // return <Message/>
  //   return <Signin />;
  return <>{appRouterSwitch("/signin")}</>;
}

export default App;
