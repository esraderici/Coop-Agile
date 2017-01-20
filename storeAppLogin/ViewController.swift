//
//  ViewController.swift
//  storeAppLogin
//
//  Created by CesurMecnun on 14/01/2017.
//  Copyright © 2017 CesurMecnun. All rights reserved.
//

import UIKit
import Alamofire
import SwiftyJSON
import SCLAlertView
import SVProgressHUD

class ViewController: UIViewController {

    @IBOutlet weak var txtEmail: UITextField!
    @IBOutlet weak var txtPassword: UITextField!
    @IBAction func fncSignIn(_ sender: UIButton) {
        
        SVProgressHUD.show(withStatus: "Loading...");
        SVProgressHUD.setDefaultMaskType(SVProgressHUDMaskType.black);
        
        let url = "http://jsonbulut.com/json/userLogin.php?ref=b7e55c8ed007bc921ac646df023a4dcd&userEmail=\(txtEmail.text!)&userPass=\(txtPassword.text!)&face=no"
        Alamofire.request(url, method: .get, encoding: URLEncoding.default).responseJSON { response in
            if let data = response.result.value {
                let json = JSON(data)
                let durum = json["user"][0]["durum"].boolValue
                let mesaj = json["user"][0]["mesaj"].string
                if durum == true {
                    let id = json["user"][0]["bilgiler"]["userId"].string
                    UserDefaults.standard.set(id!, forKey: "id");
                    SCLAlertView().showSuccess("İşlem Sonucu", subTitle: mesaj!, closeButtonTitle: "Tamam")
                }
                else {
                    SCLAlertView().showWarning("İşlem Sonucu", subTitle: mesaj!, closeButtonTitle: "Tamam")
                }
            }
            else {
                SCLAlertView().showWarning("İşlem Sonucu", subTitle: "İnternet bağlantınızı kontrol ediniz.", closeButtonTitle: "Tamam")
            }
            SVProgressHUD.dismiss();
        }
    }
    
    @IBAction func fncSignUp(_ sender: UIButton) {
        performSegue(withIdentifier: "signUp", sender: nil)
    }
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view, typically from a nib.
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }


}

