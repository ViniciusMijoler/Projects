import { toast } from 'react-toastify';

 export const information = (message: string) => {
  toast.info(message, {
    position: "bottom-right",
    autoClose: 2000,
    hideProgressBar: false,
    closeOnClick: true
  })
}

export const error = (message: string) => {
  toast.error(message, {
    position: "bottom-right",
    autoClose: 2000,
    hideProgressBar: false,
    closeOnClick: true
  })
};

export const success = (message: string) => {
  toast.success(message, {
    position: "bottom-right",
    autoClose: 2000,
    hideProgressBar: false,
    closeOnClick: true
  })
};

export const warning = (message: string) => {
  toast.warn(message, {
    position: "bottom-right",
    autoClose: 2000,
    hideProgressBar: false,
    closeOnClick: true
  })
};