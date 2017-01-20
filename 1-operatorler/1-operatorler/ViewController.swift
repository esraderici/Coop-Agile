//
//  ViewController.swift
//  1-operatorler
//
//  Created by wissen on 23/10/16.
//  Copyright © 2016 wissen. All rights reserved.
//

import UIKit

class ViewController: UIViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
        
        // arimetik operatörler
        let say1:Float = 51.6
        let say2:Float = 65.7
        var sonuc:Float = 0
        
        // toplama
        sonuc = say1 + say2
        print("Toplama Sonucu : \(sonuc)")
        
        // çıkarma
        sonuc = say1 - say2
        print("Çıkarma Sonucu : \(sonuc)")
        
        
        // bölme
        sonuc = say1 / say2
        print("Bölme Sonucu : \(sonuc)")
        
        // çarpma
        sonuc = say1 * say2
        print("Çarpma Sonucu : \(sonuc)")
        
        // Mod
        //sonuc = (say1 % say2)
        //print("Mod Sonucu : \(sonuc)")
        
        
        // artrıma ve eksiltme
        // ++, --
        var s1 = 0
        var s2 = 1
        s1 += 1
        s2 -= 1
        print("s1 \(s1)")
        print("s2 \(s2)")
        
        // s1++
        // ++s1
        
        // mantıksal operatörler
        // ==, >=,<=, ...
        // == eşitlik operatörü
        let a = 10
        let b = 20
        let durum = a == b ? 3 : 2
        print("durum değeri : \(durum)")
        
        
        // != -> eşit değil yapısı
        if (a != b) {
            print("a ve b eşit değil")
        }
        
        
        // >=, <=
        if (a >= b) {
            print("a b'den büyük veya eşittir.")
        }
        
        
        // birden fazla koşul (ve, veya) kullanımı
        // &&, ||
        // && kullanımı -> 2 koşulunda sağlanması durumu
        if (a == 10 && b == 20) {
            
        }
        
        
        // || kullanımı
        if (a == 10 || b == 20) {
            
        }
        
        
        
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }


}

