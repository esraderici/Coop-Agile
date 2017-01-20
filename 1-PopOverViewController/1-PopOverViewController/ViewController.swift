//
//  ViewController.swift
//  1-PopOverViewController
//
//  Created by CesurMecnun on 25/12/2016.
//  Copyright © 2016 CesurMecnun. All rights reserved.
//

import UIKit

class ViewController: UIViewController, UIPopoverPresentationControllerDelegate {

    @IBOutlet weak var btnGec: UIButton!
    @IBAction func fncGec(_ sender: UIButton) {
        self.performSegue(withIdentifier: "liste", sender: nil)
    }
    
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        if segue.identifier == "liste" {
            let vc = segue.destination
            vc.popoverPresentationController?.delegate = self
            vc.preferredContentSize = CGSize(width: self.view.frame.width / 2, height: 150)
            vc.popoverPresentationController?.sourceView = btnGec
            vc.popoverPresentationController?.sourceRect = btnGec.bounds
            
        }
    }
    
    func adaptivePresentationStyle(for controller: UIPresentationController) -> UIModalPresentationStyle {
        
        return .none
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

