
import { Suspense } from "react";
import { BrowserRouter } from "react-router-dom";
import { Toaster } from "react-hot-toast";
import { RouteConfig } from "./routes/config";
import GlobalCss from "./styles/global.css";


const App = () => {
    const { showPopup } = useAppSelector((state) => state.common || {});

    return (
      <BrowserRouter>
        <Toaster position="top-center" reverseOrder={false} />
        {showPopup && <SelectSizePopup />}

        
          <GlobalCss />
          <Suspense fallback={<Lottie jsonName="loading" text="loading" />}>
            <RouteConfig />
           
          </Suspense>
       
      </BrowserRouter>
    );
  };


export default App;

