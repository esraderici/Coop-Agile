//
//  ViewController.swift
//  1-TimerKullanimi
//
//  Created by Hakan on 13/11/16.
//  Copyright © 2016 Hakan. All rights reserved.
//

import UIKit

class ViewController: UIViewController {

    var time:Timer!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        time = Timer.scheduledTimer(withTimeInterval: 1.0, repeats: true){(Timer) in
            self.calis()
        }
        
    }

    
    // çalış fonksiyonu
    var i = 0
    func calis(){
        print("Çalıştı")
        i += 1
        if i == 10 {
            time.invalidate()
            self.performSegue(withIdentifier: "anaEkran", sender: nil)
        }
        
    }
    
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }


}

